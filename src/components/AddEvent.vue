<template>
<span>
    <h1>Add Event</h1>
    <v-form>
        <v-text-field label="Event Title"></v-text-field>
        <v-btn @click="toggleDatePicker">Select Date</v-btn>
        <v-btn @click="toggleTimePicker">Select Time</v-btn>
        <v-date-picker v-if="showDatePicker" @change="handleDateSelected" label="Date" v-model="date"></v-date-picker>
        <h3 v-if="date">Date: {{date}}</h3>
        <v-time-picker v-if="showTimePicker" v-model="time"></v-time-picker>
        <v-btn v-if="showTimePicker" @click="toggleTimePicker">Done</v-btn>
        <h3 v-if="time">Time: {{time}}</h3>
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
import friends from "@/data/friends.json"
export default{
    Name: "AddEvent",
    data(){
        return{
            showDatePicker: false,
            jsonFriendOptions: friends,
            friendOptions: [],
            currentFriend: "",
            friendList: [],
            date: "",
            time: "",
            showTimePicker: false
        };
    },
    methods: {
        toggleDatePicker(){
            this.showDatePicker = !this.showDatePicker;
        },
        
        addFriend(){
            if(this.friendList.indexOf(this.currentFriend) === -1){
                this.friendList.push(this.currentFriend);
            }
            this.currentFriend = "";
        },

        handleDateSelected(){
            this.toggleDatePicker();
        },

        toggleTimePicker(){
            this.showTimePicker = !this.showTimePicker;
        },

        saveEvent(){
            alert("Save Event")
        }
    },
    mounted(){
        this.jsonFriendOptions.forEach(friend => {
            let newFriend = friend.firstName + " " + friend.lastName;
            this.friendOptions.push(newFriend);
        });
    }
}
</script>