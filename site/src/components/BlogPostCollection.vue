<template>
  <div class="card-collection">
    <div class="card">
      <h2>Heading One</h2>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Qui dolorum reiciendis itaque, vitae mollitia ad
        accusantium similique pariatur minus repellat maiores voluptatibus recusandae perspiciatis fugiat provident
        nostrum assumenda inventore velit!</p>
      <img src="../../public/img/bulldogs.png" alt="Bulldog">
      <a>Read More</a>
    </div>
    <div class="card">
      <h2>Heading Two</h2>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Qui dolorum reiciendis itaque, vitae mollitia ad
        accusantium similique pariatur minus repellat maiores voluptatibus recusandae perspiciatis fugiat provident
        nostrum</p>
      <img src="../../public/img/bulldogs.png" alt="Bulldog">
      <a>Read More</a>
    </div>
    <div class="card">
      <h2>Heading Three</h2>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Qui dolorum reiciendis itaque, vitae mollitia ad
        nostrum assumenda inventore velit!</p>
      <img src="../../public/img/bulldogs.png" alt="Bulldog">
      <a>Read More</a>
    </div>
    <div class="card">
      <h2>Heading Four</h2>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Qui dolorum reiciendis itaque, vitae mollitia ad </p>
      <img src="../../public/img/bulldogs.png" alt="Bulldog">
      <a>Read More</a>
    </div>
  </div>

</template>

<style scoped>
img {
  max-width: 100%;
  display: block;
}

.card-collection {
  padding-left: 5%;
  padding-right: 5%;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  grid-template-rows: 100% repeat(auto-fit, minmax(150px, 1fr));
  gap: 1rem;
  margin-top: 5rem;
}

@media only screen and (min-width: 1200px) {
  .card-collection {
    grid-template-columns: repeat(3, minmax(350px, 1fr));
  }
}

.card {
  border: 1px solid #ccc;
  border-radius: 6px;
  padding: 2rem;
  text-align: center;
  display: grid;
  grid-template-rows: subgrid;
  grid-row: span 4;
}

.card:hover {
  box-shadow: 0px 2px 10px 0px black;
  transition: box-shadow 0.1s linear;
}
</style>
