<template>
<span>
    <h1>Choose View Section</h1>
    <calendar></calendar>
</span>
</template>

<script>
import Calendar from "./Calendar.vue"
export default{
    name: "ChooseView",
    components: {
        Calendar
    }
}
</script>