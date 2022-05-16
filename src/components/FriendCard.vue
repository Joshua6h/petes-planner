<template>
    <v-sheet class="mt-7" elevation="4" rounded="lg">
        <v-container>
            <h1>Add Friends</h1>
            <div class="ma-auto" style="width: 50%;">
                <v-text-field label="Search for a Friend" autofocus outlined v-model="friend"
                    :color="this.$store.state.accent" background-color="white"></v-text-field>
            </div>
            <v-btn type="submit" :disabled="friend == '' ? true : false">Add</v-btn>
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
    }
}
</script>