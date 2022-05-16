<template>
    
    <div>
        <h1>Calendar</h1>
        <v-sheet height="500" max-width="1000" class="ma-auto">
            <v-row>
                <v-col><v-btn @click="prev()" icon><v-icon>mdi-chevron-left</v-icon></v-btn></v-col>
                <v-col><v-select :items="['Monthly View', 'Weekly View', 'Daily View']" label="Select View" v-model="view" @change="handleViewChange"></v-select></v-col>
                <v-col><v-btn @click="next()" icon><v-icon>mdi-chevron-right</v-icon></v-btn></v-col>
            </v-row>
            <v-calendar :type="type" :events="events" :start="start" ref="calendar" outlined></v-calendar>
        </v-sheet>
    </div>
    
</template>

<script>

// import events from "@/data/events.json"
export default{
    name: "Calendar",
    data(){
        return {
            type: "month",
            // events: events,
            start: new Date(Date.now()),
            view: 'Monthly View'
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
        },
        handleViewChange(){
            if (this.view === "Monthly View"){
                this.setMonthlyView();
            }
            if (this.view === "Weekly View"){
                this.setWeeklyView();
            }
            if (this.view === "Daily View"){
                this.setDailyView();
            }
        }
    },
    async beforeMount(){
        await this.$store.dispatch('GET_EVENTS');
    },
    computed: {
        events(){
            let eventsList = this.$store.getters.events
            const formattedEvents = []
            eventsList.forEach(event => {
                let formattedEvent = {
                    id: event.event_id,
                    name: event.title,
                    description: event.description,
                    participants: event.friends,
                    start: new Date(event.start_datetime),
                    end: new Date(event.end_datetime),
                    timed: true
                };
                formattedEvents.push(formattedEvent);
             });
            return formattedEvents;
        },
    },
}
</script>