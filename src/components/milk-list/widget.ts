import { customElement, property} from 'lit/decorators.js'
import {LitElement, html, css} from 'lit'
import {Milk} from "./milk.type.ts";


@customElement('farm-milk-list')
export class FarmMilkList extends LitElement {
    static styles = css`
 
   .content {
        background: white;
        width: 80%;
        position: absolute;
        left: 10%;
        top: 50%;
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
    data: Milk[] = []

    @property({attribute: false, type: Boolean})
    isLoading = false

    private fetchData() {
        this.updateComplete.then(() => {
            this.isLoading = true

            fetch(`http://localhost:9030/milk?from=2021-01-04&to=2023-07-12`)
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

   render(){
       if (this.isLoading) {
           return html`
                <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM" crossorigin="anonymous">
                <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-geWF76RCwLtnZ8qwWowPQNguL3RmwHVBC9FhGdlKrxdiJJigb/j/68SIy3Te4Bkz" crossorigin="anonymous"></script>

                <div class="spinner-border" role="status">
                    <span class="visually-hidden">Loading...</span>
                </div>`
       }

       if (this.error !== ''){
           return html`
            <div style="    
            background: white;
            top: 50%;
            left: 40%;
            padding: 20px;
            position: absolute;
            border-radius: 10px;"> Ooops...something get wrong!</div>
            `
       }

       let rows = []

       if (this.data!=undefined){
           for (const milk of this.data) {
               let row = html`
                <tr>
                    <td>${milk.date}</td>
                    <td>${milk.liters}</td>
                    <td>${milk.price}</td>
                    <td class="table-active" >${milk.liters * milk.price}</td>
                </tr>
                `
               rows.push(row)
           }
       }

       return html`
           <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM" crossorigin="anonymous">
           <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-geWF76RCwLtnZ8qwWowPQNguL3RmwHVBC9FhGdlKrxdiJJigb/j/68SIy3Te4Bkz" crossorigin="anonymous"></script>


           <div class="content">
               <h1>Milk History</h1>
               <table class="table">
                   <thead>
                   <td>Date</td>
                   <td>Liters</td>
                   <td>PRICE/liter</td>
                   <td >Price</td>
                   </thead>
                   <tbody>
                   ${rows}
                   </tbody>
               </table>
           </div>
       `
   }


}