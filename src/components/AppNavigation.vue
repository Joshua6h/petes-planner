<template>
    <span>
        <v-toolbar elevation="3" color="#FCEDDA">
            <v-app-bar-nav-icon class="hidden-md-and-up"></v-app-bar-nav-icon>
            <router-link to="/">
                <v-app-bar-title>Pete's Planner</v-app-bar-title>
            </router-link>
            <v-toolbar flat color="#FCEDDA">
                <v-tabs slider-color="#EE4E34">
                    <v-tab to="/">About</v-tab>
                    <v-tab v-if='isAuthenticated' to="/dashboard">Dashboard</v-tab>
                    <v-tab v-if='isAuthenticated' to="/calendar">Calendar</v-tab>
                    <v-tab v-if='isAuthenticated' to="/add-events">Add Events</v-tab>
                </v-tabs>
            </v-toolbar>
            <v-spacer></v-spacer>
            <v-btn v-if='!isAuthenticated' plain outlined to="/login">Log In</v-btn>
            <div v-if='isAuthenticated && hasProfile' class="mx-3">{{getUsername}}</div>
            <v-btn v-if='isAuthenticated' to="/profile" icon> <!-- Make the sign in and account buttons alternate visibility based on authentication -->
                <v-icon>mdi-account-circle</v-icon>
            </v-btn>
            <v-btn v-if='!isAuthenticated' icon>
                <v-icon>mdi-account-circle</v-icon>
            </v-btn>
        </v-toolbar>
    </span>
</template>

<script>
export default {
    name: 'AppNavigation',
    computed: {
        isAuthenticated() {
            return this.$store.getters.isAuthenticated;
        },
        hasProfile() {
            return this.$store.getters.hasProfile;
        },
        getUsername() {
            return this.$store.getters.username;
        }
    },
    created() {
        let status = this.$store.getters.hasProfile;
        if (status == '') {
            this.$store.dispatch('GET_PROFILE');
        }
    }
}
</script>

<style scoped>
a, a:visited {
    color: black;
    text-decoration: none;
}
</style>