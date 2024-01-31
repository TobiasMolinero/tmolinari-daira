<script>
    import Calculadora from "../components/Calculadora.svelte";
    import Historial from "../components/Historial.svelte";
</script>


<div class="w-screen h-screen grid place-items-center">
    <div class="contenedor_calculadora bg-black p-2 w-72">
        <div class="header bg-gray-200">
            <h2 class="text-xl text-center font-bold py-3">Calculadora Basica</h2>
        </div>
        <div class="body bg-white p-5">
            <!-- <Calculadora /> -->
            <Historial />
        </div>
    </div>
</div>

