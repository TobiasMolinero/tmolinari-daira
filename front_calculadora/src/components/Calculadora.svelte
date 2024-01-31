<div class="container">
    <div class="pantalla_calculadora">
        <div class="pantalla_superior bg-black text-end text-white">8*</div>
        <div class="pantalla_inferior bg-black text-end text-white text-xl py-1">3</div>
    </div>
    <div class="teclado grid grid-cols-4 grid-rows-5 gap-1 mt-1">
        <button class="boton_teclado cero">CE</button>
        <button class="boton_teclado cero">C</button>
        <button class="boton_teclado operador">+/-</button>
        <button class="boton_teclado operador">%</button>
        <button class="boton_teclado number">7</button>
        <button class="boton_teclado number">8</button>
        <button class="boton_teclado number">9</button>
        <button class="boton_teclado operador">/</button>
        <button class="boton_teclado number">4</button>
        <button class="boton_teclado number">5</button>
        <button class="boton_teclado number">6</button>
        <button class="boton_teclado operador">x</button>
        <button class="boton_teclado number">1</button>
        <button class="boton_teclado number">2</button>
        <button class="boton_teclado number">3</button>
        <button class="boton_teclado operador">-</button>
        <button class="boton_teclado number">0</button>
        <button class="boton_teclado number">.</button>
        <button class="boton_teclado igual">=</button>
        <button class="boton_teclado operador">+</button>
    </div>
</div>

<style>


    .boton_teclado{
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

    .operador{
        background: maroon;
    }
    .operador:hover{
        background: rgb(95, 0, 0);
    }
    .operador:active{
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

    .igual{
        background: green;
    }
    .igual:hover{
        background: rgb(0, 85, 0);
    }
    .igual:active{
        background: green;
    }

</style>