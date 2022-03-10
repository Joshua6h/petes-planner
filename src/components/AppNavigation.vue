<template>
    <span>
        <v-toolbar elevation="3" color="#FCEDDA">
            <v-app-bar-nav-icon class="hidden-md-and-up"></v-app-bar-nav-icon>
            <router-link to="/">
                <v-app-toolbar-title >Pete's Planner</v-app-toolbar-title>
            </router-link>
            <v-toolbar flat color="#FCEDDA">
                <v-tabs slider-color="#EE4E34">
                    <v-tab to="/">Dashboard</v-tab>
                    <v-tab to="/about">About</v-tab>
                </v-tabs>
            </v-toolbar>
            <v-spacer></v-spacer>
            <v-btn plain outlined>Sign In</v-btn>
            <v-btn icon>
                <v-icon>mdi-account-circle</v-icon>
            </v-btn>
        </v-toolbar>
    </span>
</template>

<script>
export default {
    name: 'AppNavigation',
}
</script>

<style scoped>
a, a:visited {
    color: black;
    text-decoration: none;
}
</style>