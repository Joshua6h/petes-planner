<template>
    <v-sheet class="mt-7" elevation="4" rounded="lg">
        <v-container class="mb-5">
            <h1>Add Friend</h1>
            <div class="ma-auto" style="width: 50%;">
                <v-text-field label="Username" autofocus outlined v-model="friend"
                    :color="this.$store.state.accent" background-color="white"></v-text-field>
            </div>
            <v-btn @click="addFriend()" outlined color="error" :disabled="friend == '' ? true : false">Add</v-btn>
        </v-container>
        <v-container>
            <h2>Current Friends</h2>
            <v-simple-table fixed-header height="300px">
                <thead class="text-left">
                    <tr>
                        <th>Username</th>
                        <th>Name</th>
                    </tr>
                </thead>
                <tbody class="text-left">
                    <tr v-for="user in getFriends" :key="user.user_id">
                        <td>{{user.username}}</td>
                        <td>{{user.last_name}}, {{user.first_name}}</td>
                        <td class="text-right">
                            <v-btn><v-icon color="error">mdi-trash-can-outline</v-icon></v-btn>
                        </td>
                    </tr>
                </tbody>
            </v-simple-table>
        </v-container>
    </v-sheet>
</template>

<script>
export default {
    name: 'FriendCard',
    data() {
        return {
            friend: ''
        }
    },
    methods: {
        addFriend() {
            if (this.friend != '') {
                this.$store.dispatch('ADD_FRIEND', {username: this.friend});
            }
        }
    },
    computed: {
        getFriends() {
            if (this.$store.getters.getFriendsStatus === 'success') {
                return this.$store.getters.friends;
            } else {
                return [];
            }
        }
    },
    async beforeMount(){
        await this.$store.dispatch('GET_FRIENDS');
    }
}
</script>