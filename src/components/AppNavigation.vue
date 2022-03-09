<template>
    <span>
        <v-toolbar elevation="3">
            <v-app-bar-nav-icon class="hidden-md-and-up"></v-app-bar-nav-icon>
            <router-link to="/">
                <v-app-toolbar-title >Pete's Planner</v-app-toolbar-title>
            </router-link>
            <p :class="{'red--text':true}">Test text</p>
            <template>
                <v-tabs v-model="tab">
                    <v-tab to="/">Home</v-tab>
                    <v-tab to="/about">About</v-tab>
                </v-tabs>
            </template>
            <v-spacer></v-spacer>
            <v-btn plain outlined>Sign In</v-btn>
            <v-btn icon>
                <v-icon>mdi-account</v-icon>
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