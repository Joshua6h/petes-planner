<template>
<span>
    <h1>Add Event</h1>
    <v-card class="pa-4 ma-2 ma-auto" max-width="1000px">
        <v-form>
            <v-row>
                <v-col><v-text-field v-model="title" label="Event Title"></v-text-field></v-col>
                <v-col></v-col>
            </v-row>
            <v-text-field v-model="description" label="Event Description"></v-text-field>
            <v-btn @click="toggleStartDatePicker" class="ma-2">Select Start Date</v-btn>
            <v-btn @click="toggleStartTimePicker" class="ma-2">Select Start Time</v-btn><br>
            <v-date-picker v-if="showStartDatePicker" @change="handleStartDateSelected" label="Start Date" v-model="startDate"></v-date-picker>
            <h3 v-if="startDate">Start Date: {{startDate}}</h3>
            <v-time-picker v-if="showStartTimePicker" v-model="startTime"></v-time-picker>
            <v-btn v-if="showStartTimePicker" @click="toggleStartTimePicker" class="ma-2">Done</v-btn>
            <h3 v-if="startTime">Start Time: {{startTime}}</h3><br>
            <v-btn @click="toggleEndDatePicker" class="ma-2">Select End Date</v-btn>
            <v-btn @click="toggleEndTimePicker" class="ma-2">Select End Time</v-btn><br>
            <v-date-picker v-if="showEndDatePicker" @change="handleEndDateSelected" label="End Date" v-model="endDate"></v-date-picker>
            <h3 v-if="endDate">End Date: {{endDate}}</h3>
            <v-time-picker v-if="showEndTimePicker" v-model="endTime"></v-time-picker>
            <v-btn v-if="showEndTimePicker" @click="toggleEndTimePicker" class="ma-2">Done</v-btn>
            <h3 v-if="endTime">End Time: {{endTime}}</h3>
            <v-row>
                <v-col><v-select :items="friendOptions" label="Add a friend" v-model="currentFriend"></v-select></v-col>
                <v-col class="mt-2"><v-btn @click="addFriend">Add Friend</v-btn></v-col>
            </v-row>
            <div v-if="friendList.length != 0">
                <h3>Friends</h3>
                <v-list-item>
                    <v-list-item-content>
                        <v-list-item-title v-for="friend in friendList" :key="friend">{{friend}}</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
            </div>
            <v-btn @click="saveEvent">Save</v-btn>
        </v-form>
    </v-card>
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
            friendOptions: [],
            currentFriend: "",
            friendList: [],
            friendObjectList: [],
            showStartDatePicker: false,
            showStartTimePicker: false,
            startDate: "",
            startTime: "",
            showEndDatePicker: false,
            showEndTimePicker: false,
            endDate: "",
            endTime: "",
        };
    },
    methods: {
        toggleStartDatePicker(){
            this.showStartDatePicker = !this.showStartDatePicker;
        },

        toggleEndDatePicker(){
            this.showEndDatePicker = !this.showEndDatePicker;
        },
        
        addFriend(){
            if(this.friendList.indexOf(this.currentFriend) === -1){
                this.friendList.push(this.currentFriend);
                let friendObject = this.$store.getters.friends.filter(option => option.first_name + ' ' + option.last_name === this.currentFriend);
                this.friendObjectList.push(friendObject[0]);
            }
            this.currentFriend = "";
            
        },

        handleStartDateSelected(){
            this.toggleStartDatePicker();
        },

        handleEndDateSelected(){
            this.toggleEndDatePicker();
        },

        toggleStartTimePicker(){
            this.showStartTimePicker = !this.showStartTimePicker;
        },

        toggleEndTimePicker(){
            this.showEndTimePicker = !this.showEndTimePicker;
        },

        saveEvent(){
            const idList = [];
            this.friendObjectList.forEach(friend => {
                idList.push(friend.user_id.toString());
                
            })

            let event = {
                title: this.title,
                description: this.description,
                start_datetime: new Date(new Date(this.startDate + ' ' + this.startTime)),
                end_datetime: new Date(new Date(this.endDate + ' ' + this.endTime)),
                friends: idList
            };
            this.$store.dispatch("ADD_EVENT", event);
            // reset fields
            this.title = "";
            this.description = "";
            this.currentFriend = "";
            this.friendList = [];
            this.friendObjectList = [];
            this.showStartDatePicker = false;
            this.showStartTimePicker = false;
            this.startDate = "";
            this.startTime = "";
            this.showEndDatePicker = false;
            this.showEndTimePicker = false;
            this.endDate = "";
            this.endTime = "";
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