<h1>INSTRUÇÕES PARA RODAR O PROJETO</h1>

<h2>SEM DOCKER</h2>
    <h3>PRÉ-REQUISITOS:</h3>
        <ol>
            <li>Instalar Node.js: https://nodejs.org/pt</li>
            <li>Instalar Go Language: https://go.dev/</li>
            <li>Instalar pacotes React: 'npm install create-react-app'</li>
        </ol>
        <p>OBS: Usuários Linux talvez precisem instalar o pacote npm separadamente.</p>
    <h3>EXECUTANDO APLICAÇÃO: BACKEND & FRONTEND</h3>
        <ol>
            <li>Navegue até a pasta backend</li>
            <li>Execute o comando 'go run .'</li>
            <li>Para visualizar o backend: http://localhost:8080/tasks</li>
            <li>Em outro terminal, navegue até a pasta frontend</li>
            <li>Execute o comando 'npm run start'</li>
            <li>Frontend deve abrir por padrão, caso não abra: http://localhost:3000/</li>
        </ol>

<h2>COM DOCKER</h2>
    <h3>PRÉ-REQUISITOS:</h3>
        <ol>
            <li>Docker =3</li>
        </ol>
    <h3>EXECUTANDO APLICAÇÃO:</h3>
        <ol>
            <li>Abra o Docker</li>
            <li>Navegue até a raiz do projeto</li>
            <li>Execute o comando 'docker-compose up --build'</li> 
            <li>Para visualizar o backend: http://localhost:8080/tasks</li>
            <li>Para visualizar o frontend: http://localhost:3000/</li>
        </ol>



