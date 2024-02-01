<script>
// @ts-nocheck
	import Button from "./Button.svelte";

    let buttons = ["CE", "C", "+/-", "%", "7", "8", "9", "/", "4", "5", "6", "x", "1", "2", "3", "-", "0", ".", "=", "+"]
    let bottomDisplay = "";
    let topDisplay = "";
    let firstNumber = 0;
    let secondNumber = 0;
    let operator = "";
    let flag = false;
    let wasResult = false;


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
        if(value == "0" && operator == "/"){
            alert('ATENCION: no puede dividir un numero en cero');
            bottomDisplay = "";
            return false;
        }
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
    }


</script>


<div class="container">
    <div>
        <div class="bg-black text-end text-white h-6">{topDisplay}</div>
        <div class="bg-black text-end text-white text-xl py-1 h-8">{bottomDisplay}</div>
    </div>
    <div class="teclado grid grid-cols-4 grid-rows-5 gap-1 mt-1">

        {#each buttons as button}
            <Button valor={button} on:send={getValorButton} />
        {/each}

    </div>
</div>

