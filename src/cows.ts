import { customElement, property} from 'lit/decorators.js'
import { LitElement, html } from 'lit'
import {Data} from "./cow.type";

@customElement('my-page')
export class MyPage extends LitElement {
    @property({attribute: false, type: String})
    error = ''

    @property({attribute: false, type: Array})
    data: Data[]


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


    firstUpdated() {
        this.fetchData()
    }

    render() {
        let i = ""
        if (this.data!=undefined){
            i = this.data[0].Id
        }
        console.log(this.data)


        return html`
            <h1>Hello</h1>
            id: ${i}
      
         
        `}
}

//data : ${this.data[0].Id},${this.data[0].Birthday},${this.data[0].Colour},${this.data[0].MotherId}