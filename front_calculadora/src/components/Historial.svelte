<script>
// @ts-nocheck
    

	import { onMount, createEventDispatcher } from "svelte";

    let historial = [];
    let result;
    const dispatch = createEventDispatcher();

    const getHistorial = async() => {
        try {
            const response = await fetch('http://localhost:4000/historial')
            historial = await response.json();
        } catch (error) {
            alert(error);
        }
    }

    const selectOperation = (id) => {
        let operationSelected = historial.filter(h => h.id_operacion === id);
        result = operationSelected[0].resultado;
        sendResult();
    }

    const sendResult = () => {
        dispatch('send', result);
    }

    onMount(getHistorial);
</script>

<div>
    <h2 class="text-center text-4xl">Historial</h2>
    <div class="contenedor_tabla container h-72 bg-black">
        <table class="tabla text-sm bg-black text-white w-full min-h-100">
            <thead>
                <tr class="text-center text-xs">
                    <th class="w-4/12">Fecha</th>
                    <th class="w-4/12">Operacion</th>
                    <th class="w-4/12 text-center">Resultado</th>
                </tr>
            </thead>
            <tbody class="text-center text-xs">
                {#each historial as h}
                    <tr class="fila" on:click={selectOperation(h.id_operacion)} title="Click para continuar a partir de este resultado">
                        <td>{h.fecha}</td>
                        <td>{h.operacion}</td>
                        <td>{h.resultado}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>
</div>


<style>
    .contenedor_tabla{
        overflow: auto;
    }

    .fila:hover{
        background-color: gray;
        color: black;
        cursor: pointer;
    }
</style>