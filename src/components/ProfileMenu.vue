<template>
    <div>
        <v-card :color="getPrimary">
            <v-list-item-group color="error" mandatory v-model="model">
                <v-list-item v-for="(item, i) in menuItems" :key="i" @click="changePage(item.text)">
                    <v-list-item-title>{{item.text}}</v-list-item-title>
                </v-list-item>
            </v-list-item-group>
        </v-card>
    </div>
</template>

<script>
export default {
    name: 'ProfileMenu',
    data() {
        return {
            model: 0,
            menuItems: [
                {
                    text: 'PROFILE'
                },
                {
                    text: 'FRIENDS'
                },
                {
                    text: 'LOGOUT'
                }
            ]
        }
    },
    computed: {
        getPrimary() {
            return this.$store.state.primary;
        }
    },
    methods: {
        changePage(text) {
            if (text === 'LOGOUT') {
                this.logout();
            }
            else {
                this.$emit('changePage', text);
            }
        },
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