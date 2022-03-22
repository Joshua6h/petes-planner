<template>
<span>
    <h1>Add Event</h1>
    <v-form>
        <v-text-field label="Event Title"></v-text-field>
        <v-btn @click="toggleDatePicker">Select Date</v-btn>
        <v-date-picker v-if="showDatePicker" label="Date"></v-date-picker>
        <v-select :items="friendOptions" label="Add a friend" v-model="currentFriend"></v-select>
        <v-btn @click="addFriend">Add Friend</v-btn>
        <div>
            <h3>Friends</h3>
            <v-list-item>
                <v-list-item-content>
                    <v-list-item-title v-for="friend in friendList" :key="friend">{{friend}}</v-list-item-title>
                </v-list-item-content>
            </v-list-item>
        </div>
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
            friendList: []
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