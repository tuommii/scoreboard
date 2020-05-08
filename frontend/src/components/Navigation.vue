<template>
  <div class="navbar is-fixed-bottom has-text-centered">

    <div class="columns is-gapless is-mobile">
      <div class="column left"><a class="button bold" @click="navigate(showPrev)">
        Prev <span class="tag">{{showPrev}}</span></a>
      </div>
      <div class="column middle active is-2"><a class="button active">{{active}}</a></div>
      <div class="column right"><a class="button bold" @click="navigate(showNext)">
        Next <span class="tag">{{showNext}}</span></a></div>
    </div>
</div>
</template>

<script>
export default {
  props: ['active', 'basketCount'],
  data() {
    return {
    }
  },
  methods: {
    navigate(num) {
      this.$emit('navigate', num);
    }
  },
  computed: {
    showNext() {
      if (this.active >= this.basketCount)
        return (1);
      return this.active + 1;
    },
    showPrev() {
      if (this.active <= 1)
        return (this.basketCount);
      return this.active - 1;
    },
  }
};
</script>

<style scoped>
a, a:visited {
  color: #fff !important;
}

.navbar {
  min-height: 0 !important;
  background: #000 !important;
  /* overflow:hidden; */
  /* background: red !important; */
}

a.button {
  box-sizing: border-box;
  border: 0;
  background : inherit;
  width: 100%;
  height: 100%;
  padding-bottom: 0.75rem;
  padding-top: 0.75rem;
}


.active {
  padding: 0;
  margin: 0;
  vertical-align: middle;
  padding-top: 0.35rem;
}

.column {
  border: 1px solid #2a7390;
  background: #1f2743;
}

.left {
  border: 1px solid #2a7390;
  border-right: 0;
}
.right {
  border: 1px solid #2a7390;
  border-left: 0;
}
.tag {
  color: #ccc;
  background: #000;
  margin-left: 5px;
}
@media only screen and (min-width: 481px) {
  .navbar {
    /* padding: 0.75rem; */
    max-width: 480px;
    width: 100%;
    margin: 0 auto;
    display: block;
  }
  a.button {
    /* padding-bottom: 0.4rem; */
  }
}
</style>
