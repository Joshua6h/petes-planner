<template>
    <v-card>
        <v-container>
            <h2>{{username}}</h2>
            <h2>{{firstname}} {{lastname}}</h2>
            <h2 class="mb-5">{{email}}</h2>
            <!-- <h3><v-btn icon><v-icon>mdi-border-color</v-icon></v-btn> Edit Information</h3> -->
            <!-- We probably don't want them to be able to change their info at the moment -->
        </v-container>
    </v-card>
</template>

<script>
export default {
    name: 'ProfileCard',
    computed: {
        username() {
            return this.$store.getters.username;
        },
        firstname() {
            return this.$store.getters.firstname;
        },
        lastname() {
            return this.$store.getters.lastname;
        },
        email() {
            return this.$store.getters.email;
        }
    }
}
</script>