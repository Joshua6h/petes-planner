<template>
    <span>
        <h1>Welcome, Name</h1>
        <v-row class="ma-1 d-flex">
            <v-col>
                <upcoming-events></upcoming-events>
            </v-col>
            <v-col>
                <choose-view></choose-view>
                <router-link to="/calendar">Go to calendar</router-link>
            </v-col>
            <v-col>
                <add-event></add-event>
                <router-link to="/add-events">Go to add event</router-link>
            </v-col>    
        </v-row>
    </span>
</template>

<script>
import AddEvent from './AddEvent.vue';
import ChooseView from './ChooseView.vue';
import UpcomingEvents from './UpcomingEvents.vue';
export default{
  components: { UpcomingEvents, ChooseView, AddEvent },
    name: "Dashboard"
};
</script>