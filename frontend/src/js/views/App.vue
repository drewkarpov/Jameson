<template>
    <b-container class="jameson-container">
        <b-row>
            <b-col class="mt-3 mb-3">
                <h1>Image Comparator</h1>

                <router-link to="/test" class="router-link">Compare</router-link>
                <router-link to="/voidzones" class="router-link">Define voidzones</router-link>
            </b-col>
        </b-row>

        <router-view/>
    </b-container>
</template>
