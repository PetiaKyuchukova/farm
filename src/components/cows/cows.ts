import { customElement, property} from 'lit/decorators.js'
import {LitElement, html, css} from 'lit'
import {Cow} from "./cow.type.ts";

@customElement('farm-herd')
export class FarmHerd extends LitElement {
    static styles = css`
 
    .content {
        background: white;
        width: 80%;
        position: absolute;
        left: 10%;
        top: 30%;
        border-radius: 24px;
        padding: 20px;
    }
    .search {
        background: white;
        width: 80%;
        position: absolute;
        left: 10%;
        top: 20%;
        border-radius: 24px;
        padding: 20px;
    }
    h1 {
        color: #367749
    }
    `
    @property({attribute: false, type: String})
    error = ''

    @property({attribute: true, type: Boolean})
    visible = false

    @property({attribute: false, type: Array})
    data: Cow[] = []

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
        let rows = []
        if (this.data!=undefined){
            for (const cow of this.data) {
                let row = html`
                <tr>
                    <td>${cow.id}</td>
                    <td>${cow.gender}</td>
                    <td>${cow.breed}</td>
                    <td>${cow.colour}</td>
                </tr>
                `
                rows.push(row)
            }
        }


        return html`
            <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM" crossorigin="anonymous">
            <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-geWF76RCwLtnZ8qwWowPQNguL3RmwHVBC9FhGdlKrxdiJJigb/j/68SIy3Te4Bkz" crossorigin="anonymous"></script>
            
            <div class="search"> 
                <h2>Search</h2>
            </div>
            <div class="content">
                <div style="display: flex;     justify-content: space-between;">
                    <h1>Herd</h1>
                    <button type="button" style="    height: 40px;" class="btn btn-success">+ Add cow</button>
                    
                </div>
                <table class="table table-hover">
                    <thead>
                        <td>Cow ID</td>
                        <td>Gender</td>
                        <td>Breed</td>
                        <td>Color</td>
                    </thead>
                    <tbody>
                        ${rows}
                    </tbody>
                </table>
            </div>
           
        `}
}

