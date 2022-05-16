<template>
    <span>
        <h1>Upcoming Events</h1>
        <v-card class="pa-4 ma-2">
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
        </v-card>
    </span>
</template>

<script>

// import events from "@/data/events.json"
export default{
    name: "UpcomingEvents",
    data(){
        return{
            eventsLoaded: false
        };
    },
    computed: {
        events(){
            let eventsList = this.$store.getters.events
            const newEvents = []
            eventsList.forEach(event => {
                let newEvent = {
                    id: event.event_id,
                    title: event.title,
                    description: event.description,
                    friends: event.friends,
                    start_datetime: (new Date(event.start_datetime)).toLocaleDateString() + ' ' + (new Date(event.start_datetime)).toLocaleTimeString(),//((new Date(event.start_datetime)).getMonth() + 1).toString() + '/' + ((new Date(event.start_datetime)).getDate() + 1) + ' ' + ((new Date(event.start_datetime)).getHours()) + ':' + ((new Date(event.start_datetime)).getMinutes()),
                    end: new Date(event.end_datetime)
                };
                if (newEvent.end > new Date(Date.now())){
                    newEvents.push(newEvent)
                }
            });
            return newEvents
        },
    },
    async beforeMount(){
        await this.$store.dispatch('GET_EVENTS');
    }
}
</script>