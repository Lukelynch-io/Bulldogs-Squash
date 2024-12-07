<template>
  <div class="header">
    <h1><a href="#">Bulldogs Squash Club</a></h1>
    <div class="hamburger-menu">
      <a class="hamburger-menu-item" href="#">About</a>
      <a class="hamburger-menu-item" href="#">Fixtures</a>
      <a class="hamburger-menu-item" style="" href="#">Contact Us</a>
    </div>
  </div>
</template>

<style scoped>
.header {
  background-color: rgb(26 35 126 / var(--tw-bg-opacity, 1));
  border-radius: 0px 0px 5px 5px;
  width: 100%;
  padding-left: 1ch;
  display: flex;
  justify-content: space-between;
}

.header h1,
a {
  color: white;
  font-family: "Arial";
  margin-top: 0.5ch;
  margin-bottom: 0.5ch;
  text-decoration: none;
}

.hamburger-menu {
  display: table;
}

.hamburger-menu-item {
  display: table-cell;
  vertical-align: middle;
  padding-left: 2.5px;
  padding-right: 2.5px;
  font-weight: 600;
  text-decoration: none;
}

.hamburger-menu-item:last-of-type {
  border-radius: 0px 0px 5px 0px;
  padding-right: 1ch;
}

.hamburger-menu-item:hover {
  background-color: rgb(26 35 68 / var(--tw-bg-opacity, 1));
}
</style>
