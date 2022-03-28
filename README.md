# CadastroReservadeAlunos


<h3>1 - Para rodar o projeto na sua máquina</h3>
<ol>
  <li>Antes de mais nada é exencialmente nescessário que você instale o golang na sua máquina para poder rodar e compilar este projeto</li>
  <li>agora dê um git clone neste repositório</li>
<code>
  git clone https://github.com/AnDeizinho/CadastroReservadeAlunos.git
</code>
<li>Depois que tiver baixado este repositório e todos os arquivos. Localize neste repositório onde se encontra o arquivo main.go, que fica localizado no diretório raiz do 
  repositório. Em seguida abra o cmd neste diretório em questão e digite o comando a seguir</li>
<code>
  go run main.go
</code>
  <li>Para compilar o projeto digite</li>
<code>
  go build main.go
</code>
   <li>Pode ser que o Sistema Operacional solicite a autorização para o uso da porta :5000 do firewall. É só altorizar que ta tudo certo</li>
    <li>Agora é só abrir o navegador de sua preferencia e digitar "localhost:5000" no campo da url</li>
  <li>se prefirir rodar o projeto na porta :80 basta alterar a linha 17 do arquivo main.go.<br>
    <code>servidor := http.Server{Addr: ":5000", Handler: Rotas()}</code> <br>mude para <br>
    <code>servidor := http.Server{Addr: ":80", Handler: Rotas()}</code>
  </li>
</ol>

<h3>2 - Como o projeto esta estruturado</h3>
<p>
  Este projeto esta estrutura do padrão MVC. Nãe estranhe se você notar alguma semelhança com os padrões de projetos do aspnet mvc, é que eu sou charpeiro mesmo.
  
</p>
<h5>O main.go</h5>
<p>Ele que faz o milagre acontecer. Nele esta definida a função principal que instancia uma estrutura que vai ser o servidor. E define uma função Rota()
  que basicamente mapeia os arquivos do projetos.
</p>
<h5>A Pasta views</h5>
<p>
  Contem os templates que são carregados pelo servidor
</p>
<h5>A pasta controller</h5>
<p>
  Contem os Códigos que carregam os templates. Cada arquivo é um Controller, e cada Controller tem sua pastinha na views. E cada template é renderizado por funções 
  de dentro do controller. o controller tbm configura as rotas de suas respectivas views
</p>
<h5>A pasta model</h5>
  <p>
  Contem as structs responsáveis por levar os dados armazenados para o controller, e o controller leva para a views. No inverso, a views leva os dados para o controller que os trata, e passa para o model que armazena.
</p>
<h5>A pasta public</h5>
  <p>
    Contem as pastas e arquivos estáticos do projeto
</p>



