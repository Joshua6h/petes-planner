<template>
    <span>
        <v-calendar v-if="view === 'month'"></v-calendar>
    </span>
</template>

<script>
import events from "@/data/events.json"
export default{
    name: "Calendar",
    data(){
        return {
            view: "month",
            events: events
        }
    }
}
</script>