<template>
  <div class="blog-post">
    <h2>
      <slot name="blogTitle"></slot>
    </h2>
    <p>
      <slot name="blogDescription"></slot>
    </p>
  </div>
</template>

<style scoped>
.blog-post {
  border: 2px solid black;
  /* Black border */
  padding: 20px;
  /* Add padding inside the border */
  margin: 20px 0;
  /* Space between blog posts */
  border-radius: 8px;
  /* Rounded corners */
  background-color: #f9f9f9;
  /* Light gray background */
}

.blog-post h2 {
  margin-bottom: 10px;
  /* Space between title and description */
  text-align: center;
}

.blog-post p {
  margin-top: 0;
  /* Remove top margin from paragraph */
}
</style>
