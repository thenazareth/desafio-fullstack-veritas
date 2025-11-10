<h1>INSTRUÇÕES PARA RODAR O PROJETO: WINDOWS (e linux)</h1>

<h2>SEM DOCKER</h2>
    <h3>PRÉ-REQUISITOS:</h3>
        <ol>
            <li>Instalar Node.js: https://nodejs.org/pt</li>
            <li>Instalar Go Language: https://go.dev/</li>
            <li>Instalar pacotes React: 'npm install create-react-app'</li>
        </ol>
        <p>Usuários Linux talvez precisem instalar o pacote npm separadamente.</p>
    <h3>EXECUTANDO APLICAÇÃO: BACKEND & FRONTEND</h3>
        <ol>
            <li>Navegue até a pasta backend</li>
            <li>Execute o comando 'go run .'</li>
            <li>Para visualizar o backend: <a href="http://localhost:8080/tasks">http://localhost:8080/tasks</a></li>
            <li>Em outro terminal, navegue até a pasta frontend</li>
            <li>Execute o comando 'npm run start'</li>
            <li>Frontend deve abrir por padrão, caso não abra: <a href="http://localhost:3000/">http://localhost:3000/</a></li>
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
            <li>Para visualizar o backend: <a href="http://localhost:8080/tasks">http://localhost:8080/tasks</a></li>
            <li>Para visualizar o frontend: <a href="http://localhost:3000/">http://localhost:3000/</a></li>
        </ol>
    <p>Usuários de Linux podem seguir os seguintes comandos:</p>
    <ul>
        <li>sudo apt-get docker (baixar docker)</li>
        <li>sudo apt-get docker-compose (puglin docker-compose)</li>
        <li>sudo systemctl start docker (inicar servidor manualmente)</li>
        <li>docker compose up --build (executar a aplicação)</li>
    </ul>
    <p>Para visualizar o backend: <a href="http://localhost:8080/tasks">http://localhost:8080/tasks</a></p>
    <p>Para visualizar o frontend: <a href="http://localhost:3000/">http://localhost:3000/</a></p>
    



