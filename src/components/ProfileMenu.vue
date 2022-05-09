<template>
    <div class="pa-3">
        <v-btn class="menu-button" tile depressed color="error" outlined>Friends</v-btn>
        <v-btn class="menu-button" tile depressed color="error" outlined>Events</v-btn>
        <v-btn class="menu-button" tile depressed color="error" v-on:click="logout">Logout</v-btn>
    </div>
</template>

<script>
export default {
    name: 'ProfileMenu',
    methods: {
        logout: function() {
            this.$store.dispatch('LOGOUT').then(() => {
                this.$router.push('/login');
            }).catch((err) => {
                if (err.response) {
                    // Request made and server responded
                    this.$emit('raiseError', err.response.data);
                }
            })
        }
    }
}
</script>

<style scoped>
.menu-button {
    width: 100%;
}
</style>