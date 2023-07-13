import {customElement, property,state} from "lit/decorators.js";
import {LitElement,html, nothing, PropertyValues} from "lit";
import {Cow, Insemination, Pregnancy} from "../cows/cow.type.ts";

@customElement('farm-cow-profile')
export class FarmCowProfile extends LitElement {
    @property({reflect: true, attribute: 'cow'})
    private cow: string

    @property({reflect: true, attribute: 'visible'})
    private visible: string

    @state()
    private visibleB = false

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

    @state()
    emptyCow: Cow = {
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

    @property({attribute: false, type: Boolean})
    isLoading = false

    @property({attribute: false, type: String})
    error = ''


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

    private closeCowProfile(e) {
            this.visibleB = false
    }

    private saveCowProfile(e) {
        fetch(`http://localhost:9030/upsert`, {
            method: 'PUT',
            body: JSON.stringify(this.data)
        }).then(async (response) => {
            if (response.ok) {
                //const savedBudget = await response.json()
                console.log("Saved!")
            } else {
                this.error = 'Error saving cow.'
                console.log("err saving cow!")
            }
        })

        this.visibleB = false
    }

    handleVisibility(){
        this.fetchData()

        if (this.cow ==  ''){
            this.visibleB = false
            return
        }
       this.visibleB = true
    }



    updated(changedProperties: PropertyValues) {
        const hasOrgChanged = changedProperties.has('visible')
        if (hasOrgChanged ) {
            this.handleVisibility()
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
                <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM" crossorigin="anonymous">
                <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-geWF76RCwLtnZ8qwWowPQNguL3RmwHVBC9FhGdlKrxdiJJigb/j/68SIy3Te4Bkz" crossorigin="anonymous"></script>
                
                <div class="content" style="    
                background: #fff;
                width: 69%;
                position: absolute;
                top: 15%;
                left: 15%;
                padding: 20px;
                border-radius: 10px;">
                    <div style="    
                    background: #ced08866;
                    padding: 20px;
                    border-radius: 10px;
margin-bottom: 20px;">
                        <h1 style=" color: #3f7c4b;">${this.data.id} Profile</h1>
                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Birthday</span>
                            <input type="date" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm">
                        </div>
                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Gender</span>
                            <input type="text" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm">
                        </div>
                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Color</span>
                            <input type="text" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm" value="${this.data.colour}" onchange="${this.data.colour = 'changed'}">
                        </div>
                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Breed</span>
                            <input type="text" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm">
                        </div>
                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Mother ID</span>
                            <input type="text" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm">
                        </div>
                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Father ID</span>
                            <input type="text" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm">
                        </div>
                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Father Breed</span>
                            <input type="text" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm">
                        </div>
                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Is Pregnant</span>
                            <input type="text" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm" value="${this.data.isPregnant}">
                        </div>

                        <div class="input-group input-group-sm mb-3">
                            <span class="input-group-text" id="inputGroup-sizing-sm">Last Ovulation</span>
                            <input type="date" class="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-sm" value="${this.data.ovulation}">
                        </div>
                    </div>
                    <div style="    
                    background: #ced08866;
                    padding: 20px;
                    border-radius: 10px;
margin-bottom: 20px;">
                        <h3>Inseminations</h3>
                        
                    </div>
                    <div style="    
                    background: #ced08866;
                    padding: 20px;
                    border-radius: 10px;
margin-bottom: 20px;">
                        <h3>Pregnancies</h3>
                    </div>
                    <div>
                        <button @click="${this.closeCowProfile}">Cancel</button>
                        <button @click="${this.saveCowProfile}">Save</button>
                    </div>
                </div>
                
             
            ` : nothing

        return profile
    }
}