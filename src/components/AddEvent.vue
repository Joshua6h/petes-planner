<template>
<span>
    <h1>Add Event</h1>
    <v-form>
        <v-text-field v-model="title" label="Event Title"></v-text-field>
        <v-text-field v-model="description" label="Event Description"></v-text-field>
        <v-btn @click="toggleStartDatePicker">Select Date</v-btn>
        <v-btn @click="toggleStartTimePicker">Select Time</v-btn>
        <v-date-picker v-if="showStartDatePicker" @change="handleStartDateSelected" label="Start Date" v-model="startDate"></v-date-picker>
        <h3 v-if="startDate">Start Date: {{startDate}}</h3>
        <v-time-picker v-if="showStartTimePicker" v-model="startTime"></v-time-picker>
        <v-btn v-if="showStartTimePicker" @click="toggleStartTimePicker">Done</v-btn>
        <h3 v-if="startTime">Start Time: {{startTime}}</h3>
        <v-select :items="friendOptions" label="Add a friend" v-model="currentFriend"></v-select>
        <v-btn @click="addFriend">Add Friend</v-btn>
        <div>
            <h3 v-if="friendList.length != 0">Friends</h3>
            <v-list-item>
                <v-list-item-content>
                    <v-list-item-title v-for="friend in friendList" :key="friend">{{friend}}</v-list-item-title>
                </v-list-item-content>
            </v-list-item>
        </div>
        <v-btn @click="saveEvent">Save</v-btn>
    </v-form>
</span>
</template>

<script>
// import friends from "@/data/friends.json"
export default{
    Name: "AddEvent",
    data(){
        return{
            title: "",
            description: "",
            showStartDatePicker: false,
            // jsonFriendOptions: friends,
            friendOptions: [],
            currentFriend: "",
            friendList: [],
            startDate: "",
            startTime: "",
            showStartTimePicker: false
        };
    },
    methods: {
        toggleStartDatePicker(){
            this.showStartDatePicker = !this.showStartDatePicker;
        },
        
        addFriend(){
            if(this.friendList.indexOf(this.currentFriend) === -1){
                this.friendList.push(this.currentFriend);
            }
            this.currentFriend = "";
        },

        handleStartDateSelected(){
            this.toggleStartDatePicker();
        },

        toggleStartTimePicker(){
            this.showStartTimePicker = !this.showStartTimePicker;
        },

        saveEvent(){
            alert("Save Event")
            // let event = {
            //     title: this.title,
            //     description: this.description,
            //     start_datetime: this.date + ' ' + this.time
            // }
        }
    },
    async beforeMount(){
        await this.$store.dispatch('GET_FRIENDS');
        this.$store.getters.friends.forEach(friend => {
            let newFriend = friend.first_name + " " + friend.last_name;
            this.friendOptions.push(newFriend);
        });
    },
}
</script>