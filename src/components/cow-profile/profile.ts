import {customElement, property, state} from "lit/decorators.js";
import {LitElement, html, nothing, PropertyValues} from "lit";
import {Cow, Insemination, Pregnancy} from "../cows/cow.type.ts";


@customElement('farm-cow-profile')
export class FarmCowProfile extends LitElement {
    @property({reflect: true, attribute: 'cow'})
    private cow: string

    @property({reflect: true, attribute: 'visible'})
    private visible: string

    @state()
    private visibleB = false

    @state()
    private visibleDeletion = false

    @state()
    private visibleInseminations = false

    @state()
    private addingInseminations = false

    @state()
    private addingPregnancy = false

    @state()
    private visiblePregnancies = false

    @property({attribute: false, type: Object})
    data: Cow = {
        id: "",
        colour: "",
        birthdate: new Date(),
        gender: "",
        breed: "",
        ovulation: new Date(),
        motherId: "",
        farmerId: "",
        fatherBreed: "",
        inseminations: [],
        pregnancies: [],
        isPregnant: false
    }

    @property({attribute: false, type: Object})
    addedInsemination: Insemination = {
        date: new Date(),
        breed: "",
        IsArtificial: false,
    }

    @property({attribute: false, type: Object})
    addedPregnancy: Pregnancy = {
        detectedAt: new Date('0001-01-01'),
        firstDay: new Date('0001-01-01'),
        lastDay: new Date('0001-01-01')
    }

    @property({attribute: false, type: Boolean})
    isLoading = false

    @property({attribute: false, type: String})
    error = ''

    getToday(){
        const date = new Date();

        let day = date.getDate();
        let month = date.getMonth() + 1;
        let year = date.getFullYear();

        let currentDate = `${day}-${month}-${year}`;

        return currentDate
    }

    private fetchData() {
        this.updateComplete.then(() => {
            this.isLoading = true

            fetch(`http://localhost:9030/cows/${this.cow}`)
                .then(async resp => {
                    console.log(this.data)
                    if (resp.status === 200) {
                        this.data = await resp.json()
                    } else {
                        this.error = 'Error loading cow.'
                    }
                    this.isLoading = false
                })
                .catch(reason => {
                    this.error = `Network error: ${reason}`
                    this.isLoading = false
                })
        })
    }
    closeDeletForm(){
        this.visibleDeletion = false
    }

    private closeCowProfile() {
        this.visibleB = false
        this.addedPregnancy =  {
            detectedAt: new Date('0001-01-01'),
            firstDay: new Date('0001-01-01'),
            lastDay: new Date('0001-01-01')
        }
        this.addedInsemination = {
            date: new Date(0 ,0,0),
            breed: "",
            IsArtificial: false,
        }
        this.data  = {
            id: "",
            colour: "",
            birthdate: new Date(),
            gender: "",
            breed: "",
            ovulation: new Date(),
            motherId: "",
            motherBreed: "",
            fatherId: "",
            fatherBreed: "",
            inseminations: [],
            pregnancies: [],
            isPregnant: false
        }
    }
    private pregnanciesVisibility() {
        this.visiblePregnancies = !this.visiblePregnancies
        this.addingPregnancy = false
    }
    private inseminationsVisibility() {
        this.visibleInseminations = !this.visibleInseminations
        this.addingInseminations = false
    }

    onChangeColor(e:any) { this.data.colour = e.target.value }
    onChangeID(e:any) { this.data.id = e.target.value }
    onChangeBirthdate(e:any) { this.data.birthdate = (e.target.value)}
    onChangeGender(e:any) { this.data.gender = e.target.value }
    onChangeBreed(e:any) { this.data.breed = e.target.value }
    onChangeMotherID(e:any) { this.data.motherId = e.target.value }
    onChangeMotherBreed(e:any) { this.data.motherBreed = e.target.value }
    onChangeFatherID(e:any) { this.data.fatherId = e.target.value }
    onChangeFatherBreed(e:any) { this.data.fatherBreed = e.target.value }
    onChangePregnancy(e:any) {
        if (e.target.checked){
            this.data.isPregnant = true
        }}
    onChangeNotPregnancy(e:any) {
        if (e.target.checked){
         this.data.isPregnant = false
    } }
    onChangeOvulation(e:any) {
        this.data.ovulation = e.target.value
    }
    onChangeLastPregnancyLastDay( idx: number){
         return (e: any) => {
             this.data.pregnancies[idx].lastDay = e.target.value
         }
    }
    onChangeLastPregnancyFirstDay( idx: number){
         return (e: any) => {
             this.data.pregnancies[idx].firstDay = e.target.value
         }
    }
    onChangeAddedPregnancyDetection(e:any) { this.addedPregnancy.detectedAt = e.target.value }
    onChangeAddedPregnancyFirstDay(e:any) { this.addedPregnancy.firstDay = e.target.value }
    onChangeAddedPregnancyLastDay(e:any) { this.addedPregnancy.lastDay = e.target.value }
    onChangeAddedInseminationDate(e:any) { this.addedInsemination.date = e.target.value }
    onChangeAddedInseminationBreed(e:any) { this.addedInsemination.breed = e.target.value }
    onChangeAddedInseminationIsArtf(e:any) { this.addedInsemination.IsArtificial = e.target.value }


    private saveCowProfile() {
        if (this.addingPregnancy){
            this.data.pregnancies.push(this.addedPregnancy)
        }
        if (this.addingInseminations){
            this.data.inseminations.push(this.addedInsemination)
        }
        console.log("save")
        fetch(`http://localhost:9030/upsert`, {
            method: 'PUT',
            body: JSON.stringify(this.data)
        }).then(async (response) => {
            if (response.ok) {
                console.log("Saved!")
            } else {
                this.error = 'Error saving cow.'
                console.log("err saving cow!")
            }
        })

            this.visibleB = false
            this.addedPregnancy =  {
                detectedAt: new Date('0001-01-01'),
                firstDay: new Date('0001-01-01'),
                lastDay: new Date('0001-01-01')
            }
            this.addedInsemination = {
                date: new Date(0 ,0,0),
                breed: "",
                IsArtificial: false,
            }
            window.location.reload();
            return
    }


     deleteCow() {
        console.log(this.cow)
        fetch(`http://localhost:9030/delete/${this.cow}`, {
            method: 'DELETE',
        }).then(async (response) => {
            if (response.ok) {
                console.log("Deleted!")
            } else {
                this.error = 'Error deleting cow.'
                console.log("err deleting cow!")
            }
        })

         this.visibleB = false
         this.addedPregnancy =  {
             detectedAt: new Date('0001-01-01'),
             firstDay: new Date('0001-01-01'),
             lastDay: new Date('0001-01-01')
         }
         this.addedInsemination = {
             date: new Date(0 ,0,0),
             breed: "",
             IsArtificial: false,
         }
         window.location.reload();
         return
    }

    confirmDeletion() {
        this.visibleDeletion = true
        console.log(this.visibleDeletion )

    }

    handleVisibility(){
        if(this.cow != 'new'){
            this.fetchData()
        }

        if (this.cow ==  ''){
            this.visibleB = false
            return
        }
       this.visibleB = true
    }

    addInseminations(){
        this.addingInseminations = !this.addingInseminations
        this.visibleInseminations = true
    }
    addPregnancy(){
        this.addingPregnancy = !this.addingPregnancy
        this.visiblePregnancies = true
    }

    renderInseminations(){
        let rows = []

        for (const insemination of this.data.inseminations) {
            let row = html`
            <tr>
                <td>${insemination.date}</td>
                <td>${insemination.breed}</td>
                <td>${insemination.IsArtificial}</td>
            </tr>
            `
            rows.push(row)
        }


        let additionalRow = this.addingInseminations? html`
            <tr style="opacity: ">
                <td>
                    <div class="input-group input-group-sm mb-3">
                    <input type="date" id="insemination-date" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm" @change="${this.onChangeAddedInseminationDate}" >
                    </div>
                </td>
                <td>
                    <div class="input-group input-group-sm mb-3">
                        <input type="text" id="insemination-breed" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm" @change="${this.onChangeAddedInseminationBreed}" >
                    </div>
                </td>
                <td>
                    <div class="input-group input-group-sm mb-3">
                        <input type="text" id="insemination-artf" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm" @change="${this.onChangeAddedInseminationIsArtf}">
                    </div>
                </td>
            </tr>` : nothing

        rows.push(additionalRow)


        return html`
            <table style="width: 100%">
                <thead>
                <tr>
                    <th scope="col">Date</th>
                    <th scope="col">Breed</th>
                    <th scope="col">Is Artificial?</th>
                </tr>
                </thead>
                ${rows}
            </table>`


    }
    renderPregnancies(){
        let rows = []

        for (let i = 0; i < this.data.pregnancies.length; i++) {
            if (i === this.data.pregnancies.length - 1 && this.data.isPregnant){
                    let row = html`
                        <tr>
                            <td>
                                ${this.data.pregnancies[i].detectedAt}
                            </td>
                            <td>
                                <div class="input-group input-group-sm mb-3">
                                    <input type="date" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm" value="${this.data.pregnancies[i].firstDay}" @change="${this.onChangeLastPregnancyFirstDay(i)}">
                                </div>
                            </td>
                            <td>
                                <div class="input-group input-group-sm mb-3">
                                    <input type="date" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm" value="${this.data.pregnancies[i].lastDay}" @change="${this.onChangeLastPregnancyLastDay(i)}" >
                                </div>
                            </td>
                        </tr>
                `
                rows.push(row)
                console.log(row)
                continue
            }
            let row = html`
            <tr>
                <td>${this.data.pregnancies[i].detectedAt}</td>
                <td>${this.data.pregnancies[i].firstDay}</td>
                <td>${this.data.pregnancies[i].lastDay}</td>
            </tr>
            `
            rows.push(row)
        }

        let additionalRow = this.addingPregnancy? html`
            <tr>
                <td>
                    <div class="input-group input-group-sm mb-3">
                    <input type="date" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm" @change="${this.onChangeAddedPregnancyDetection}">
                    </div>
                </td>
                <td>
                    <div class="input-group input-group-sm mb-3">
                        <input type="date" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm" @change="${this.onChangeAddedPregnancyFirstDay}" >
                    </div>
                </td>
                <td>
                    <div class="input-group input-group-sm mb-3">
                        <input type="date" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm" @change="${this.onChangeAddedPregnancyLastDay}">
                    </div>
                </td>
            </tr>` : nothing

        rows.push(additionalRow)

        return html`<table style="width: 100%; ">
            <thead>
            <tr>
                <th scope="col">Detected at</th>
                <th scope="col">First Day</th>
                <th scope="col">Last Day</th>
            </tr>
            </thead>
            ${rows}
        </table>`
    }

    updated(changedProperties: PropertyValues) {
        const hasVisibleChanged = changedProperties.has('visible')
        if (hasVisibleChanged) {
            this.handleVisibility()
            this.visiblePregnancies = false
            this.visibleInseminations = false
        }
    }

    render() {

        if (this.isLoading) {
            return html`
                <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM" crossorigin="anonymous">
                <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-geWF76RCwLtnZ8qwWowPQNguL3RmwHVBC9FhGdlKrxdiJJigb/j/68SIy3Te4Bkz" crossorigin="anonymous"></script>

                <div class="spinner-border" role="status">
                    <span class="visually-hidden">Loading...</span>
                </div>`
        }
        let profile = this.visibleB ?
            html`
                <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet"
                      integrity="sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM"
                      crossorigin="anonymous">
                <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"
                        integrity="sha384-geWF76RCwLtnZ8qwWowPQNguL3RmwHVBC9FhGdlKrxdiJJigb/j/68SIy3Te4Bkz"
                        crossorigin="anonymous"></script>

                <div class="content" style="    
                background: #fff;
                width: 69%;
                position: absolute;
                top: 8%;
                left: 15%;
                padding: 20px;
                border-radius: 10px;
    box-shadow: 1px 1px 20px 12px gray;">
                    <div style="    background: #ced08866;padding: 20px; border-radius: 10px;margin-bottom: 20px;">
                        <h1 style=" color: #3f7c4b;">${this.data.id} Profile</h1>
                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Id</span>
                            <input type="text" id="cow" class="form-control" aria-label="Sizing example input"
                                   aria-describedby="inputGroup-sizing-sm" value="${this.data.id}"
                                   @change="${this.onChangeID}">
                        </div>
                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Birthday</span>
                            <input type="date" id="birthday" class="form-control" aria-label="Sizing example input"
                                   aria-describedby="inputGroup-sizing-sm" value="${this.data.birthdate}"
                                   @change="${this.onChangeBirthdate}">
                        </div>
                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Gender</span>
                            <input type="text" id="gender" class="form-control" aria-label="Sizing example input"
                                   aria-describedby="inputGroup-sizing-sm" value="${this.data.gender}"
                                   @change="${this.onChangeGender}">
                        </div>
                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Color</span>
                            <input type="text" id="color" class="form-control" aria-label="Sizing example input"
                                   aria-describedby="inputGroup-sizing-sm" value="${this.data.colour}"
                                   @change="${this.onChangeColor}">
                        </div>
                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Breed</span>
                            <input type="text" id="breed" class="form-control" aria-label="Sizing example input"
                                   aria-describedby="inputGroup-sizing-sm" value="${this.data.breed}"
                                   @change="${this.onChangeBreed}">
                        </div>
                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Mother ID</span>
                            <input type="text" id="motherId" class="form-control" aria-label="Sizing example input"
                                   aria-describedby="inputGroup-sizing-sm" value="${this.data.motherId}"
                                   @change="${this.onChangeMotherID}">
                        </div>
                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Mother Breed</span>
                            <input type="text" id="motherId" class="form-control" aria-label="Sizing example input"
                                   aria-describedby="inputGroup-sizing-sm" value="${this.data.motherBreed}"
                                   @change="${this.onChangeMotherBreed}">
                        </div>
                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Father ID</span>
                            <input type="text" id="fatherId" class="form-control" aria-label="Sizing example input"
                                   aria-describedby="inputGroup-sizing-sm" value="${this.data.fatherId}"
                                   @change="${this.onChangeFatherID}">
                        </div>
                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Father Breed</span>
                            <input type="text" id="fatherBreed" class="form-control" aria-label="Sizing example input"
                                   aria-describedby="inputGroup-sizing-sm" value="${this.data.fatherBreed}"
                                   @change="${this.onChangeFatherBreed}">
                        </div>
                        ${this.data.gender == "male" ? nothing : 
                                html`
                                    <div class="input-group input-group-sm mb-3">
                                        <span class="input-group-text" id="inputGroup-sizing-sm">Is Pregnant</span>
                                        <input type="radio" class="btn-check"
                                               name="options-outlined" id="success-outlined" autocomplete="on"
                                               @change="${this.onChangePregnancy}">
                                        <label class="btn btn-outline-success" for="success-outlined">Pregnant</label>

                                        <input type="radio" class="btn-check" name="options-outlined" id="danger-outlined"
                                               autocomplete="off" @change="${this.onChangeNotPregnancy}">
                                        <label class="btn btn-outline-danger" checked for="danger-outlined">NOT Pregnant</label>
                                    </div>

                                    <div class="input-group input-group-sm mb-3">
                                        <span class="input-group-text" id="inputGroup-sizing-sm">Last Ovulation</span>
                                        <input type="date" id="ovulation" class="form-control" aria-label="Sizing example input"
                                               aria-describedby="inputGroup-sizing-sm" value="${this.data.ovulation}"
                                               @change="${this.onChangeOvulation}">
                                    </div>
                                `}
                        
                    </div>
                    ${this.data.gender == "male" ? nothing : html`
                    <div style="    
                    background: #ced08866;
                    padding: 20px;
                    border-radius: 10px;
margin-bottom: 20px;">
                        <div style="display: flex;
    justify-content: space-between;">
                            <h3 @click="${this.inseminationsVisibility}" style="cursor:pointer">Inseminations</h3>
                            <button style="width: 45px" @click="${this.addInseminations}" type="button" class="btn btn-success">
                                ${this.addingInseminations && this.visibleInseminations ? "-" : "+"}
                            </button>
                        </div>

                        ${this.visibleInseminations ?
                            this.renderInseminations()
                            : nothing}
                    </div>
                    <div style="    
                    background: #ced08866;
                    padding: 20px;
                    border-radius: 10px;
margin-bottom: 20px;">
                        <div style="display: flex;justify-content: space-between;">
                            <h3 @click="${this.pregnanciesVisibility}" style="cursor:pointer">Pregnancies</h3>
                            <button style="width: 45px" @click="${this.addPregnancy}" type="button" class="btn btn-success">
                                ${this.addingPregnancy && this.visiblePregnancies ? "-" : "+"}
                            </button>
                        </div>
                        ${this.visiblePregnancies ?
                            this.renderPregnancies()
                            : nothing}
                    </div>
                    `}
                    
                    <div style="    display: flex;
    justify-content: space-between;">
                        <button class="btn btn-danger"  @click="${this.confirmDeletion}">Delete</button>
                        <div>
                            <button style="width: 100px ;" class="btn btn-outline-secondary" @click="${this.closeCowProfile}">Cancel</button>
                            <button style="width: 100px ;" class="btn btn-success" @click="${this.saveCowProfile}">Save</button>
                        </div>
                    </div>
                    
                    ${this.visibleDeletion ? html`
                        <div style="   box-shadow: 1px 1px 20px 12px gray; top: 30%; left: 22%; width: 50%; position: absolute; background: #dc3545; height: 20vh;border-radius: 10px; padding: 20px;">
                            <div style="background: white; border-radius: 10px; text-align: center;     height: 100%; padding: 20px">
                                <h4 style="margin-bottom: 25px;">
                                    Deletion of cow <strong>${this.cow}</strong>.<br>
                                    Are you sure?
                                </h4>
                                <div style="display: flex; justify-content: space-between;">
                                    <button class="btn btn-outline-secondary" @click="${this.closeDeletForm}">Cancel</button>
                                    <button class="btn btn-danger" @click="${this.deleteCow}">Delete</button>
                                </div>
                            </div>
                        </div>` : nothing}
                   
                </div>
            ` : nothing

        return profile
    }
}