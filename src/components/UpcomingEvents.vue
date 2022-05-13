<template>
    <span>
        <h1>Upcoming Events</h1>
        <v-simple-table>
            <tr>
                <th>Date</th>
                <th>Event</th>
                <th>Participants</th>
            </tr>
            <tr v-for="event in events" :key="event.event_id">
                <td>{{event.start_datetime}}</td>
                <td class="pa-5">{{event.title}}</td>
                <td class="pa-5"><v-row v-for="friend in event.friends" :key="friend">{{friend}}</v-row></td>
            </tr>
        </v-simple-table>
    </span>
</template>

<script>

// import events from "@/data/events.json"
export default{
    name: "UpcomingEvents",
    data(){
        return{
            // events: events
            // events: this.$store.getters.events
        };
    },
    computed: {
        events(){
            let eventsList = this.$store.getters.events
            const newEvents = []
            eventsList.forEach(event => {
                let newEvent = event;
                let temp = new Date(event.start_datetime);
                newEvent.start_datetime = (temp.getMonth() + 1).toString() + '/' + (temp.getDate()).toString()
                newEvent.start_datetime = newEvent.start_datetime + ' ' + temp.getHours().toString() + ':' + temp.getMinutes().toString();
                newEvents.push(newEvent)
            });
            return eventsList
        },
    },
    async beforeMount(){
        await this.$store.dispatch('GET_EVENTS');
    }
}
</script>