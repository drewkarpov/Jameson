<template>
    <span :class="className">
        {{ value }} {{ postfix }}
    </span>
</template>


<script>
export default {
    name: 'Percentage',
    props: ['value'],
    computed: {
        className: function () {
            if (this.value === '???') { // @here you'll probably want to adjust color thresholds
                return 'percentage-unknown';
            } else if (this.value >= 99) {
                return 'percentage-awesome';
            } else if (this.value >= 95) {
                return 'percentage-good';
            } else if (this.value >= 70) {
                return 'percentage-bad';
            }

            return 'percentage-awful';
        },
        postfix: function () {
            if (this.value === '???') {
                return '';
            }

            if (this.value === 100) {
                return '% 🎉🎉🎉';
            }

            return '%';
        }
    },
}

</script>
