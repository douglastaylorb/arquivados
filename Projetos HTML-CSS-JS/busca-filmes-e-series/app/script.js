async function search() {
    const searchInput = document.getElementById('searchInput').value;
    const typeSelect = document.getElementById('typeSelect').value;
    const streamingSelect = document.getElementById('streamingSelect').value;

    const url = `https://streaming-availability.p.rapidapi.com/search/filters?services=${streamingSelect}&country=br&keyword=${searchInput}&output_language=en&order_by=original_title&genres_relation=and&show_type=${typeSelect}`;
    const options = {
        method: 'GET',
        headers: {
            'X-RapidAPI-Key': '79af55c7a4msh82e2726c0de6996p1185e3jsn05a38b9394c2',
            // 'X-RapidAPI-Key': 'fc516a8c36msh409d2db1da2b750p123f1djsn298436206966',
            'X-RapidAPI-Host': 'streaming-availability.p.rapidapi.com'
        }
    };

    try {
        const response = await fetch(url, options);
        const data = await response.json();
        console.log('Dados recebidos:', data);
        if (data && data.result && data.result.length > 0) {
            await displayResults(data.result);
        } else {
            console.log('Nenhum resultado encontrado.');
            displayResults([]);
        }
    } catch (error) {
        console.error(error);
    }
}

async function displayResults(results) {
    const resultsDiv = document.getElementById('results');
    resultsDiv.innerHTML = '';

    if (results && results.length > 0) {
        const gridContainer = document.createElement('div');
        gridContainer.classList.add('grid-container');

        for (const item of results) {
            const gridItem = document.createElement('div');
            gridItem.classList.add('grid-item');

            // Chama a função para buscar o poster
            const posterUrl = await searchForPoster(item.title, item.type);
            
            // Cria um elemento img e define a URL do poster como src
            const poster = document.createElement('img');
            poster.src = posterUrl;
            poster.alt = item.title;
            poster.addEventListener('click', () => openModal(item));
            
            // Adiciona o elemento img ao gridItem
            gridItem.appendChild(poster);
            
            // Adiciona o gridItem ao gridContainer
            gridContainer.appendChild(gridItem);
        }

        resultsDiv.appendChild(gridContainer);
    } else {
        resultsDiv.innerHTML = '<div>Nenhum resultado encontrado.</div>';
    }
}

async function searchForPoster(title, type) {
    const apiKey = '4b8f1fd2';
    const url = `https://www.omdbapi.com/?apikey=${apiKey}&t=${encodeURIComponent(title)}&type=${type}`;

    try {
        const response = await fetch(url);
        const data = await response.json();
        if (data.Poster && data.Poster !== 'N/A') {
            return data.Poster;
        } else {
            return 'url_para_poster_padrao'; // Se nenhum poster for encontrado, você pode retornar a URL de um poster padrão ou null
        }
    } catch (error) {
        console.error('Erro ao buscar poster:', error);
        return 'url_para_poster_padrao';
    }
}

async function openModal(item) {
    const modalContent = document.getElementById('modalContent');

    let genresString = '';
    if (item.genres && item.genres.length > 0) {
        genresString = item.genres.map(genre => genre.name).join(', ');
    }

    if (item.type === 'movie') {
        modalContent.innerHTML = `
        <h2>${item.title}</h2>
        <ul>
            <li class="border-up"><strong>Ano:</strong> ${item.year}</li>
            <li id="gray-hover"><strong>Gêneros:</strong> ${genresString}</li>
            <li><strong>Streaming:</strong> ${item.streamingInfo?.br[0]?.service}</li>
            <li id="gray-hover"><strong>Sinopse:</strong> ${item.overview}</li>
            <li><strong>Atores:</strong> ${item.cast.join(', ')}</li>
            <li class="border-bottom" id="gray-hover"><strong>Link para acessar o filme: </strong><a href="${item.streamingInfo?.br[0]?.link}" target="_blank">Clique aqui!</a></li>
        </ul>
        `;
    } else if (item.type === 'series') {
        modalContent.innerHTML = `
            <h2>${item.title}</h2>
            <ul>
                <li class="border-up" id="gray-hover"><strong>Ano de lançamento:</strong> ${item.firstAirYear}</li>
                <li><strong>Temporadas:</strong> ${item.seasonCount}</li>
                <li id="gray-hover"><strong>Gêneros:</strong> ${genresString}</li>
                <li><strong>Streaming:</strong> ${item.streamingInfo?.br[0]?.service}</li>
                <li id="gray-hover"><strong>Sinopse:</strong> ${item.overview}</li>
                <li><strong>Principais atores:</strong> ${item.cast.join(', ')}</li>
                <li class="border-bottom" id="gray-hover"><strong>Link para acessar a série: </strong><a href="${item.streamingInfo?.br[0]?.link}" target="_blank">Clique aqui!</a></li>
            </ul>
        `;
    }

    // Abre o modal
    const modal = document.getElementById('modal');
    modal.style.display = 'block';
}

// Fecha o modal quando o usuário clica fora do conteúdo
window.onclick = function(event) {
    const modal = document.getElementById('modal');
    if (event.target === modal) {
        modal.style.display = 'none';
    }
}
