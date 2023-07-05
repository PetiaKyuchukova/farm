import { customElement, property} from 'lit/decorators.js'
import { LitElement, html } from 'lit'
import {Cow} from "./cow.type.ts";

@customElement('farm-herd')
export class FarmHerd extends LitElement {
    @property({attribute: false, type: String})
    error = ''

    @property({attribute: true, type: Boolean})
    visible = false

    @property({attribute: false, type: Array})
    data: Cow[]


    @property({attribute: false, type: Boolean})
    isLoading = false

    private fetchData() {
        this.updateComplete.then(() => {
            this.isLoading = true

            fetch(`http://localhost:9030/cows`)
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

    private redirectTo(e) {
        this.visible = false
    }


    firstUpdated() {
        this.fetchData()
    }

    render() {
        let i = "55"
        if (this.data!=undefined){
            i = this.data[0].id
        }
        console.log("from cows.", this.data)


        return html`
            <div ?hidden=${!this.visible}>

                <h1>Hello</h1>
                <button @click=${this.redirectTo}>close</button>
                id: ${i}
            </div>
           
        `}
}

//data : ${this.data[0].Id},${this.data[0].Birthday},${this.data[0].Colour},${this.data[0].MotherId}