<h1 align="center">
  <img alt="Clean Architecture" src="./public/clean.jpg" />
  <br>
  Clean Architecture 
</h1>

<h3 align="center">
  Responsabilidades de cada diretório
</h3>

<h4>
  <p>src</p>
  <ul>
    <li>
      <span>
        `domain` -> Pasta onde irá centralizar a regra de negócio do software
      </span>
    </li>
    <li>
      <span>
        `application` -> Pasta onde irá aplicar a regra de negócio do software
      </span>
    </li>
    <li>
      <span>
        `adapters` -> Pasta onde irá conter os controllers da aplicação
      </span>
    </li>
    <li>
      <span>
        `drivers` -> Pasta onde irá conter as libs instaladas no projeto
      </span>
    </li>
  </ul>
  <p>A pasta `drivers` só poderá ter interação com a pasta `adapters`.</p>
  <p>A pasta `adapters` ela poderá interagir com a `application`</p>
  <p>A pasta `application` ela poderá interagir com a `domain`</p>
  <p>A pasta `domain` ela não irá se comunicar com nenhum dos diretórios</p>
</h4>
