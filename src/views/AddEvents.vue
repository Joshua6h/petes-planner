<template>
  <div>
    <add-event></add-event>
  </div>
</template>

<script>
import AddEvent from '../components/AddEvent.vue'
// @ is an alias to /src

export default {
  name: 'Home',
  components: {
    AddEvent
  }
}
</script>