<template>
    <span class="ma-1">
        <v-btn @click="setMonthlyView" class="ma-1">Monthly View</v-btn>
        <v-spacer></v-spacer>
        <v-btn @click="setWeeklyView" class="ma-1">Weekly View</v-btn>
        <v-spacer></v-spacer>
        <v-btn @click="setDailyView" class="ma-1">Daily View</v-btn>
        <v-spacer></v-spacer>
        <v-btn @click="prev()" icon><v-icon>mdi-chevron-left</v-icon></v-btn>
        <v-btn @click="$refs.calendar.next()" value="2022-01-01" icon><v-icon>mdi-chevron-right</v-icon></v-btn>
        <v-calendar :type="type" :events="events" ref="calendar"></v-calendar>
    </span>
</template>

<script>
import events from "@/data/events.json"
export default{
    name: "Calendar",
    data(){
        return {
            type: "month",
            events: events,
            // validMonths: ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"],
            // startDate: Date.now()
        }
    },
    methods: {
        setMonthlyView(){
            this.type = 'month';
        },
        setWeeklyView(){
            this.type = 'week';
        },
        setDailyView(){
            this.type = 'day';
        },
        prev(){
            console.log(this.$refs.calendar.prev());
            this.$refs.calendar.prev();
        }
    }
}
</script>