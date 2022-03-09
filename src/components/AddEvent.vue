<template>
<span>
    <h1>Add Event</h1>
    <v-form>
        <v-text-field label="Event Title"></v-text-field>
        <v-btn>Select Date</v-btn>
        <v-date-picker label="Date"></v-date-picker>
    </v-form>
</span>
</template>

<script>
export default{
    Name: "AddEvent",
    data(){
        return{

        };
    },
    methods: {

    }
}
</script>