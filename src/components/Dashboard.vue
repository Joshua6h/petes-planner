<template>
    <span>
        <h1 class="mt-4">Welcome, {{name}}</h1>
        <v-row class="ma-1 d-flex">
            <v-col class="ma-3">
                
                <add-event class="ma-4"></add-event>
                <router-link to="/add-events" class="ma-4">Go to add event</router-link>
                <upcoming-events class="ma-4"></upcoming-events>
            </v-col>
            <v-col class="ma-3">
                <v-sheet height="600">
                    <choose-view></choose-view>
                </v-sheet> 
                <router-link to="/calendar">Go to calendar</router-link>
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
    name: "Dashboard",
    data() {
        return{
            name: this.$store.getters.firstname
        }
    }
};
</script>