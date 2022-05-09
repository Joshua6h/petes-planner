<template>
    <span class="ma-1">
        <v-btn @click="setMonthlyView" class="ma-1">Monthly View</v-btn>
        <v-spacer></v-spacer>
        <v-btn @click="setWeeklyView" class="ma-1">Weekly View</v-btn>
        <v-spacer></v-spacer>
        <v-btn @click="setDailyView" class="ma-1">Daily View</v-btn>
        <v-spacer></v-spacer>
        <v-btn @click="prev()" icon><v-icon>mdi-chevron-left</v-icon></v-btn>
        <v-btn @click="next()" icon><v-icon>mdi-chevron-right</v-icon></v-btn>
        <v-calendar :type="type" :events="events" :start="start" ref="calendar"></v-calendar>
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
            start: new Date(Date.now()),
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
            let day = this.start.getDate();
            let month = this.start.getMonth();
            let year = this.start.getFullYear();
            if(this.type === "month"){
                let myDate = new Date(year, month - 1, day);
                this.start = myDate;
            }
            if(this.type === "week"){
                let myDate = new Date(year, month, day - 7);
                this.start = myDate;
            }
            if(this.type === "day"){
                let myDate = new Date(year, month, day - 1);
                this.start = myDate;
            }
        },
        next(){
            let day = this.start.getDate();
            let month = this.start.getMonth();
            let year = this.start.getFullYear();
            if(this.type === "month"){
                let myDate = new Date(year, month + 1, day);
                this.start = myDate;
            }
            if(this.type === "week"){
                let myDate = new Date(year, month, day + 7);
                this.start = myDate;
            }
            if(this.type === "day"){
                let myDate = new Date(year, month, day + 1);
                this.start = myDate;
            }
        }
    }
}
</script>