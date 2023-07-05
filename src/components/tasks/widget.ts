import { customElement, property, state} from 'lit/decorators.js'
import { LitElement, html, nothing } from 'lit'
import {Task} from "./task.type.ts";
import "../cows/cows.ts"

@customElement('farm-tasks')
export class FarmTasks extends LitElement {



    @property({attribute: false, type: String})
    error = ''

    @property({attribute: false, type: Array})
    data: Task[]

    @state()
    private visible = false

    @property({attribute: false, type: Boolean})
    isLoading = false

    private fetchData() {
        this.updateComplete.then(() => {
            this.isLoading = true

            fetch(`http://localhost:9030/tasks?date=2023-07-02`)
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
        this.visible = true
    }
    firstUpdated() {
        this.fetchData()
    }
    render() {

        let i = "55"
        if (this.data!=undefined){
            i = this.data[0].cow_id
        }
        console.log(this.data)


        return html`
            <farm-herd id="herd" visible="true"></farm-herd>
            <h1>Tasks</h1>
            <button @click=${this.redirectTo}>cow profile</button>
            
            
        `}

}