<template>
    <span>
        <h1>Upcoming Events</h1>
        <v-simple-table>
            <tr>
                <th>Date</th>
                <th>Event</th>
                <th>Participants</th>
            </tr>
            <tr v-for="event in events" :key="event.id">
                <td>{{event.datetime}}</td>
                <td>{{event.name}}</td>
                <td><v-row><v-col v-for="participant in event.participants" :key="participant">{{participant}}</v-col></v-row></td>
            </tr>
        </v-simple-table>
    </span>
</template>

<script>
import events from "@/data/events.json"
export default{
    name: "UpcomingEvents",
    data(){
        return{
            events: events
        };
    }
}
</script>