<script>
// @ts-nocheck
	import Button from "./Button.svelte";
    import { browser } from "$app/environment";

    let buttons = ["CE", "C", "+/-", "%", "7", "8", "9", "/", "4", "5", "6", "x", "1", "2", "3", "-", "0", ".", "=", "+"]
    export let data;
    let bottomDisplay = data.result || "";
    let topDisplay = "";
    let firstNumber = 0;
    let secondNumber = 0;
    let operator = "";
    let operation = ""
    let idOperation = browser && JSON.parse(sessionStorage.getItem('id')) || 0;
    let flag = false;
    let wasResult = data.wasresult || false;
    let result = 0;


    const getValorButton = (e) => {
        let valueButton = e.detail
        if(/^[0-9]$/.test(valueButton)){
            writeBottomDisplay(valueButton);
        } else if(valueButton == '/' || valueButton == 'x' || valueButton == '-' || valueButton == '+'){
            if(flag){
                return false;
            }
            writeTopDisplay(valueButton)
        } else if(valueButton == 'CE' || valueButton == 'C'){
            if(valueButton == 'C'){
                clearDisplay();
            } else {
                bottomDisplay = "";
            }   
        } else if(valueButton == '='){
            if(!flag){
                return false;
            }
            resolve(bottomDisplay);
        }
    }

    const writeBottomDisplay = (value) => {
        if(wasResult){
            bottomDisplay = value;
            wasResult = false;
            return false;
        } else if(bottomDisplay == "0"){
            bottomDisplay = value;
            return false;
        } else if(bottomDisplay.length === 2){
            return false
        }
        bottomDisplay += value;
    }

    const writeTopDisplay = (value) => {
        if(bottomDisplay == ""){
            return false;
        }
        let sign = value == 'x' ? '*' : value;
        operator = value;
        firstNumber = +bottomDisplay;
        topDisplay = bottomDisplay + sign;
        bottomDisplay = "";
        flag = true;
    }

    const clearDisplay = () => {
        bottomDisplay = "";
        topDisplay = "";
        operator = "";
        flag = false;
    }

    const resolve = (value) => {
        const today = new Date();
        const date = today.toLocaleDateString();

        if(value == "0" && operator == "/"){
            alert('ATENCION: no puede dividir un numero en cero');
            bottomDisplay = "";
            return false;
        }
        operation = firstNumber + operator + value;
        secondNumber = +value;
        switch (operator) {
            case "/":
                bottomDisplay = firstNumber / secondNumber;
                topDisplay = "";
                flag = false;
                wasResult = true;
                break;

            case "x":
                bottomDisplay = firstNumber * secondNumber;
                topDisplay = "";
                flag = false;
                wasResult = true;
                break;

            case "-":
                bottomDisplay = firstNumber - secondNumber;
                topDisplay = "";
                flag = false;
                wasResult = true;
                break;

            case "+":
                bottomDisplay = firstNumber + secondNumber;
                topDisplay = "";
                flag = false;
                wasResult = true;
                break;
        
            default:
                break;
        }
        result = bottomDisplay.toString();
        idOperation += 1;
        browser && sessionStorage.setItem('id', idOperation);
        fetch('http://localhost:4000/historial/add', {
            method: 'POST',
            body: JSON.stringify({
                id_operacion: idOperation,
                fecha: date,
                operacion: operation,
                resultado: result
            })
        })
        .then(res => {})
        .catch(err => {alert(err)})
        ;
    }


</script>


<div class="container">
    <div>
        <div class="bg-black text-end text-white h-6">{topDisplay}</div>
        <div class="bg-black text-end text-white text-xl py-1 h-8">{bottomDisplay === "" ? "0" : bottomDisplay}</div>
    </div>
    <div class="teclado grid grid-cols-4 grid-rows-5 gap-1 mt-1">

        {#each buttons as button}
            <Button valor={button} on:send={getValorButton} />
        {/each}

    </div>
</div>

