<script>
// @ts-nocheck
    import {createEventDispatcher} from 'svelte';
    export let valor;

    const dispatch = createEventDispatcher();

    const sendValueButton = (e) => {
        // Se dispara un evento personalizado que pasa el valor del boton.
        dispatch('send', e.target.innerHTML);
    }   

</script>

{#if valor === 'CE' || valor === 'C'}

    <button class='cero' on:click={sendValueButton}>{valor}</button>

{:else if valor === '+/-' || valor === '%' || valor === '/' || valor === 'x' || valor === '-' || valor === '+'}

    <button class='operator' on:click={sendValueButton} >{valor}</button>

{:else if /^[0-9]$/.test(valor) || valor == '.'} 

    <button class="number" on:click={sendValueButton} >{valor}</button>

{:else}
    
    <button class="equal" on:click={sendValueButton}>{valor}</button>

{/if}

<!-- disabled={flag ? true : false} -->
<style>

    button{
        padding: 15px;
        color: white;
        font-weight: 700;
    }

    .cero{
        background: orangered;
    }
    .cero:hover{
        background: rgb(200, 50, 0);
    }
    .cero:active{
        background: orangered;
    }

    .operator{
        background: maroon;
    }
    .operator:hover{
        background: rgb(95, 0, 0);
    }
    .operator:active{
        background: maroon;
    }

    .number{
        background: blue;
    }
    .number:hover{
        background: rgb(0, 0, 175);
    }
    .number:active{
        background: blue;
    }

    .equal{
        background: green;
    }
    .equal:hover{
        background: rgb(0, 85, 0);
    }
    .equal:active{
        background: green;
    }

</style>