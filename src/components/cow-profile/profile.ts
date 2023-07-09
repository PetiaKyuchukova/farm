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
                        this.error = 'Error loading cost anomalies.'
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
        let profile = this.visibleB ?
            html`
                <h1>${this.data.id} Profile</h1>
                <button @click="${this.closeCowProfile}">close</button>
            ` : nothing

        return profile
    }
}