<template>
  <div class="home">
    <dashboard></dashboard>
    <!-- <HelloWorld msg="Welcome to Pete's Planner"/> -->
  </div>
</template>

<script>
// @ is an alias to /src
import Dashboard from '../components/Dashboard.vue'

export default {
  name: 'Home',
  components: {
    Dashboard
  }
}
</script>
