<template>
    <span>
        <h1 class="mt-4">Welcome, {{getName}}</h1>
        <v-row class="ma-1 d-flex">
            <v-col class="ma-3">
                <add-event class="ma-4"></add-event><br>
                <router-link to="/add-events" class="ma-4">Go to add event</router-link><br>
                <upcoming-events class="ma-4"></upcoming-events>
            </v-col>
            <v-col class="ma-3">
                <choose-view></choose-view>
                <v-spacer></v-spacer><br><br><br><br>
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
    computed: {
        getName() {
            if (this.$store.getters.hasProfile) {
                let first = this.$store.getters.firstname;
                let last = this.$store.getters.lastname;
                return first + ' ' + last;
            } else {
                return '';
            }
        }
    }
};
</script>