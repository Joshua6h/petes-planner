<template>
  <div>
    <calendar-view></calendar-view>
  </div>
</template>

<script>
import CalendarView from '../components/Calendar.vue'
// @ is an alias to /src

export default {
  name: 'Calendar',
  components: {
    CalendarView
  }
}
</script>