function tocaSom (seletorAudio) {
    const elemento = document.querySelector(seletorAudio);
    
    if(elemento === null) {
        console.log('Elemento não existe');
    } else {
        const todosElementosDeAudio = document.querySelectorAll('audio');
        todosElementosDeAudio.forEach(audio => audio.pause());
        
        elemento.currentTime = 0;
        elemento.play();
    }

}

function pausaSom() {
    const todosElementosDeAudio = document.querySelectorAll('audio');
    todosElementosDeAudio.forEach(audio => audio.pause());
}

const listaDeTeclas = document.querySelectorAll('.tecla');


for (let contador = 0; contador < listaDeTeclas.length; contador++) {

    const tecla = listaDeTeclas[contador];
    const instrumento = tecla.classList[1];

    const idAudio = `#som_${instrumento}`; // template string

    tecla.onclick = function(){
        tocaSom(idAudio);
    };

    tecla.onkeydown = function(event){
        
        if(event.code === 'Space' || event.code === 'Enter' ) {
            tecla.classList.add('ativa');
        }

    }

    tecla.onkeyup = function(){
        tecla.classList.remove('ativa');
    }
}

