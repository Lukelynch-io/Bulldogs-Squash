<template>
  <div class="header">
    <div class="flex-1"></div>
    <h1>
      <RouterLink to="/">Bulldogs Squash Club</RouterLink>
    </h1>
    <div class="flex flex-1">
      <ul class="hamburger-menu">
        <li>
          <a class="hamburger-menu-item" href="#">About</a>
        </li>
        <li>
          <a class="hamburger-menu-item" href="#">Fixtures</a>
        </li>
        <li>
          <a class="hamburger-menu-item" style="" href="#">Contact Us</a>
        </li>
        <li>
          <RouterLink class="hamburger-menu-item" to="/Login">Login</RouterLink>
        </li>
      </ul>
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
  align-items: center;
  justify-content: flex-end;
}

.flex {
  display: flex;
  justify-content: flex-end;
}

.flex-1 {
  flex: 1;
}

/* TODO: Clean this up*/
.header h1,
a {
  color: white;
  font-family: "Arial";
  margin-top: 0.5ch;
  margin-bottom: 0.5ch;
  text-decoration: none;
}

.hamburger-menu {
  display: flex;
  list-style: none;
  justify-content: space-between;
}

.hamburger-menu-item {
  padding-left: 2.5px;
  padding-right: 2.5px;
  font-weight: 600;
  text-decoration: none;
  white-space: nowrap;
  overflow: hidden;
}

.hamburger-menu-item:last-of-type {
  border-radius: 0px 0px 5px 0px;
  padding-right: 1ch;
}

.hamburger-menu-item:hover {
  text-decoration-line: underline;
}
</style>
