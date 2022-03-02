<template>
    <span>
        <h1>Welcome, Name</h1>
        <v-row>
            <v-col>
                <upcoming-events></upcoming-events>
            </v-col>
            <v-col>
                <choose-view></choose-view>
            </v-col>
            <v-col>
                <add-event></add-event>
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