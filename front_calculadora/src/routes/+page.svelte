<script>
// @ts-nocheck

    import Calculadora from "../components/Calculadora.svelte";
    import Historial from "../components/Historial.svelte";

    let historial = false;
    let data = {};

    const mostrarHistorial = () => {
        historial ? historial = false : historial = true;
    }

    const getResult = (e) => {
        data = {
            result: e.detail,
            wasresult: true
        }
        historial = false;
    }


</script>


<div class="w-screen h-screen flex justify-center items-center gap-4">
    <div class="contenedor_calculadora bg-black p-2 w-72">
        <div class="header bg-gray-200">
            <h2 class="text-xl text-center font-bold py-3">Calculadora Basica</h2>
        </div>
        <div class="body bg-white p-5">
            {#if historial}
                <Historial on:send={getResult}/>
            {:else}
                <Calculadora data={data}/>
            {/if}
        </div>
    </div>
    <button on:click={mostrarHistorial} class="bg-sky-500 py-3 px-5 rounded-md shadow-md shadow-gray-500 text-white font-semibold">
        Ver {historial ? "Calculadora" : "Historial"}
    </button>
</div>


<style>
    button:hover{
        background-color: rgb(11, 128, 182);
    }
    button:active{
        background-color: rgb(14 165 233);
        box-shadow: none;
    }
</style>
