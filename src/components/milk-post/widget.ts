import { customElement, query,property,state} from 'lit/decorators.js'
import {LitElement, html, css, nothing} from 'lit'
import {Milk} from "../milk-list/milk.type.ts";
import {PropertyValues} from "lit/development";


@customElement('farm-milk-post')
export class FarmMilkPost extends LitElement {
    static styles = css`
 
.content {
           background: white;
    width: 50%;
    width: ;
    position: absolute;
    left: 25%;
    top: 30%;
    border-radius: 24px;
    padding: 20px;
    box-shadow: 1px 1px 20px 12px gray;
    }
    h1 {
        color: #367749
    }
    `


    @property({reflect: true, attribute: 'visible'})
    private visible: string

    @state()
    private visibleB = false

    @property({attribute: false, type: Boolean})
    isLoading = false

    @property({attribute: false, type: String})
    error = ''

    @property({attribute: false, type: Object})
    data: Milk = {
        date: new Date(),
        liters:0,
        price:0
    }

    @query("#totalPrice")
    totalPrice: HTMLElement

    handleVisibility(){
        if (this.visible ==  ''){
            this.visibleB = false
            return
        }
        this.visibleB = true
    }

    updated(changedProperties: PropertyValues) {
        const hasVisibleChanged = changedProperties.has('visible')
        if (hasVisibleChanged) {
            this.handleVisibility()
        }
    }

    onChangeMilkDate(e:any) { this.data.date = (e.target.value)}
    onChangeMilkLiters(e:any) { this.data.liters = parseFloat(e.target.value)}
    onChangeMilkPrice(e:any) { this.data.price = parseFloat(e.target.value)
       let total = this.data.liters *this.data.price

             this.totalPrice.setAttribute("value", total.toString())
    }
    private saveMilk() {

        fetch(`http://localhost:9030/milk`, {
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

        this.data = {
            date: new Date(),
            liters:0,
            price:0
        }

        this.visibleB = false
        window.location.reload();
    }

    render(){
        let form = this.visibleB ?
         html`
            <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM" crossorigin="anonymous">
            <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-geWF76RCwLtnZ8qwWowPQNguL3RmwHVBC9FhGdlKrxdiJJigb/j/68SIy3Te4Bkz" crossorigin="anonymous"></script>

            
            <div class="content"> 
                <h1>Milk Today</h1>

                <div class="input-group input-group-sm mb-3">
                    <span class="input-group-text" id="inputGroup-sizing-sm">Date</span>
                    <input type="date" id="birthday" class="form-control" aria-label="Sizing example input"
                           aria-describedby="inputGroup-sizing-sm"
                           @change="${this.onChangeMilkDate}">
                </div>
                <div class="input-group input-group-sm mb-3">
                    <span class="input-group-text" id="inputGroup-sizing-sm">Liters</span>
                    <input type="number" id="gender" class="form-control" aria-label="Sizing example input"
                           aria-describedby="inputGroup-sizing-sm"
                           @change="${this.onChangeMilkLiters}">
                </div>
                <div class="input-group input-group-sm mb-3">
                    <span class="input-group-text" id="inputGroup-sizing-sm">Price/liter</span>
                    <input type="number" id="color" class="form-control" aria-label="Sizing example input"
                           aria-describedby="inputGroup-sizing-sm" 
                           @change="${this.onChangeMilkPrice}">
                </div>
                <hr>
                <div class="input-group input-group-sm mb-3">
                    <span class="input-group-text" id="inputGroup-sizing-sm">Total price</span>
                    <input type="number" id="totalPrice" class="form-control" aria-label="Sizing example input"
                           aria-describedby="inputGroup-sizing-sm" 
                           value="${this.totalPrice} lv"
                </div>
                <div>
                    <button type="button" style="    height: 40px;" class="btn btn-success" @click="${this.saveMilk}">Save</button>
                </div>
                
            </div>
        ` : nothing

        return form
    }
}