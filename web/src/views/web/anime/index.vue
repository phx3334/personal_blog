<template>
  <div class="anime-page">
    <web-navbar :no-scroll="true"/>
    <div class="container">
      <br>
      <div class="anime-sections">
        <div v-for="section in animeSections" :key="section.type" class="anime-section">
          <h2 class="section-title">{{ section.title }}</h2>
          <div class="anime-grid">
            <div 
              v-for="anime in section.animes" 
              :key="anime.id" 
              class="anime-card"
            >
              <div class="anime-card-inner">
                <img :src="anime.image" :alt="anime.title" class="anime-image">
                <div class="anime-overlay">
                  <div class="anime-info">
                    <h3 class="anime-title">{{ anime.title }}</h3>
                    <p class="anime-description">{{ anime.description }}</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import WebNavbar from '@/components/layout/WebNavbar.vue'

interface Anime {
  id: number
  title: string
  image: string
  description: string
}

interface AnimeSection {
  type: string
  title: string
  animes: Anime[]
}

const animeSections = ref<AnimeSection[]>([
  {
    type: 'light bai he',
    title: '轻百合',
    animes: [
      {id: 0, title: '魔女之旅', image: "/image/魔女之旅.jpg", description: '《魔女之旅》讲述天才少女魔女伊蕾娜，为实现冒险梦想周游列国，在一个个奇幻国度邂逅不同人与故事，在相遇与离别中，以旅人身份独自前行的公路番故事。'},
      { id: 1, title: '超时空辉耀姬', image: "/image/超时空辉耀姬.jpg", description: "《超时空辉耀姬!》(2026)是 Netflix 原创动画电影，改编自《竹取物语》，以现代东京与虚拟空间 “月读” 为舞台，讲述少女酒寄彩叶从发光电线杆捡到神秘婴儿，其飞速成长为少女后自称辉耀姬，二人以音乐为纽带在虚拟世界展开羁绊与命运交织的冒险，穿插虚拟主播、竞技与追逃剧情，画风绚丽、音乐由 VOCALOID 全明星操刀，是科幻与奇幻结合的音乐动画。" },
      { id: 2, title: '莉可莉丝', image: "/image/莉可莉丝.jpg", description: "近未来东京，秘密组织 DA 训练少女特工 “莉可丽丝” 于暗中维护和平。史上最强的千束与高冷的泷奈，在伪装成咖啡馆的分部执行任务，在日常与战斗中结下深刻羁绊。" }
    ]
  },
  {
    type: 'gentleman',
    title: '绅士区',
    animes: [
      { id: 3, title: '棍勇', image: '/image/烧火棍.jpg', description: "棍之勇者凯亚尔遭勇者团残酷虐待，凭借能力回溯时间，保留记忆与痛苦，向公主、剑勇等仇人展开疯狂复仇，改写悲惨命运的热血复仇番。" },
      { id: 4, title: '出包王女第一季', image: 'https://image.huyajs.com/cover/09ffc64839b15450d1239229d9e1c7a6.jpg', description: "纯情高中生结城梨斗向暗恋的春菜表白时，意外与外星公主菈菈结缘并开始同居，因菈菈的发明与外星追兵，开启充满意外与搞笑的后宫日常。" },
      { id: 5, title: '出包王女第二季', image: 'https://image.huyajs.com/cover/2dfbe1c802c78a8b1f75ef13b70c1124.jpg', description: '延续第一季故事，梨斗与菈菈、春菜的三角关系持续升温，金色暗影、古手川唯等更多角色登场，搞笑福利与意外日常不断升级。' },
      { id: 6, title: '出包王女第三季', image: 'https://image.huyajs.com/cover/91270e97491a918496e6569c9fcff808.jpg', description: '剧情重心转向菈菈的妹妹梦梦，她为让梨斗与所有喜欢他的女孩幸福，启动 “后宫计划”，同时杀手芽亚登场，主线与福利全面升级。' },
      { id: 7, title: '出包王女第四季', image: 'https://image.huyajs.com/cover/3bbd8afd89f7b6895b3d201b67f6ff12.jpg', description: '梦梦的后宫计划进入白热化，梨斗与众多少女的关系愈发复杂，最后讨厌H的小暗堕入黑暗后将剧情推向高潮。'}
    ]
  },
  {
    type: 'hotblood',
    title: '热血类',
    animes: [
      { id: 8, title: '咒术回战第一季', image: 'https://ts2.tc.mm.bing.net/th/id/OIP-C.dYhC8UC8-XxPTulu_cXncgAAAA?rs=1&pid=ImgDetMain&o=7&rm=3', description: '人类负面情绪化生咒灵，最强体育生虎杖悠仁吞下宿傩手指成为宿主，加入咒术高专，与伏黑惠、钉崎野蔷薇一同学习、祓除咒灵，对抗诅咒世界。' },
      { id: 9, title: '咒术回战第二季', image: 'https://img.picbf.com/upload/vod/20251207-1/d0c2c966dbe51867b72465f2cd344394.webp', description: '分为怀玉・玉折与涩谷事变两篇。前者讲述五条悟与夏油杰的高专过去以及与伏黑甚尔的激烈战斗；后者为惨烈大战，五条悟被封印，东京沦为咒灵战场。' },
      { id: 10, title: '咒术回战第三季', image: 'https://img.jisuimage.com/cover/2f9e2b5cf0b130e6316aa43850581086.jpg', description: '涩谷事变后，羂索发动 “死灭回游” 生存游戏，强制术师互相残杀。虎杖、乙骨忧太等人入局，为解封五条悟、拯救世界而战，其中伏黑惠和邮票哥的大象压房子以及乙骨的仙台一串四尤为激烈' },
      { id: 11, title: '咒术会战之纯爱战神', image: 'https://img.jisuimage.com/cover/5f221f5f69bcd62696d967a8e84e450f.jpg', description: '特级咒术师乙骨忧太被诅咒束缚，与祈本里香的纯爱羁绊成为力量核心，为守护里香与世界，以 “纯爱” 之力对抗强敌夏油杰。' },
      { id: 12, title: '进击的巨人', image: 'https://image.huyajs.com/cover/3f16df3165aeb4df25222b804214ca97.jpg', description: '人类为躲避巨人困于高墙内，艾伦目睹母亲被吃后加入兵团，逐渐揭开巨人起源、世界真相，最终走向地鸣灭世与自由抉择的悲剧史诗。' },
      { id: 13, title: '地狱乐', image: 'https://img.jisuimage.com/cover/ff5552bd9e89226e299c74531dc90590.jpg', description: '江户时代，最强忍者画眉丸被判死刑却无法被杀死。幕府令其与处刑人佐切前往 “神仙乡” 寻找不死仙药，以换自由，在岛上与怪物、死囚展开死斗。' },
      { id: 14, title: '咔哒咔哒', image: 'https://pic.youkupic.com/upload/vod/20250706-1/4448f327905f036518fcaa16dcab6877.jpg', description: '等级森严的世界里，贫民窟少年路德被诬陷杀害养父，坠入充满怪物的深渊 “奈落”。为复仇与真相，他加入清道夫，觉醒力量对抗斑兽。' },
      { id: 15, title: 'jojo的奇妙冒险', image: 'https://image.huyajs.com/cover/ba9c4fb8d7cd19c4322660138a28a427.jpg', description: '第一部《幻影血脉》讲述英国贵族乔纳森・乔斯达与心怀野心的迪奥・布兰度从小对立，后者戴上石鬼面化为吸血鬼，乔纳森修行波纹气功与之对抗，最终以同归于尽的方式结束了初代宿命。' },
      { id: 16, title: 'jojo的奇妙冒险第二季', image: '/image/龙舌兰姑娘.jpg', description: '第二部《战斗潮流》转向乔纳森之孙乔瑟夫・乔斯达，他凭借机智与波纹联手伙伴对抗苏醒的上古柱之男，阻止了他们成为究极生物毁灭世界。' },
      { id: 17, title: 'jojo的奇妙冒险之星尘斗士', image: '/image/星尘斗士.jpg', description: '第三部《星尘斗士》开启替身时代，空条承太郎为解救被 DIO 诅咒的母亲，与乔瑟夫等人横跨世界寻找并最终击败复活的 DIO，成为系列经典篇章。' },
      { id: 18, title: 'jojo的奇妙冒险之不灭钻石', image: 'https://img.jisuimage.com/cover/5b629f02e23fcc3ef15e60dcf3dd8fcf.jpg', description: '第四部《不灭钻石》舞台来到平静的杜王町，乔瑟夫私生子东方仗助与一众替身使者共同追查连环杀手吉良吉影，守护小镇日常' },
      { id: 19, title: 'jojo的奇妙冒险之黄金之风', image: '/image/黄金之风.jpg', description: '第五部《黄金之风》聚焦意大利，DIO 之子乔鲁诺・乔巴拿立志成为流氓巨星，加入黑帮组织并带领小队反抗残暴老板，为正义与梦想奋战' },
      { id: 20, title: 'jojo的奇妙冒险之石之海', image: 'https://image.huyajs.com/cover/10becf01d442fc47e00f8d7737cebc86.jpg', description: '第六部《石之海》以空条承太郎之女空条徐伦为主角，她身陷监狱，集结伙伴对抗 DIO 旧部普奇神父，在一场关乎世界重置的决战中赌上一切。' },
      { id: 21, title: 'jojo的奇妙冒险之彪马野郎', image: 'https://image.huyajs.com/cover/b8bf98401c8ed957704fc86e6d33b2a5.jpg', description: '第七部《飙马野郎》开启平行世界，残疾骑手乔尼・乔斯达与杰洛・齐贝林为争夺圣人遗体参加横跨美洲的赛马大赛，在旅途中彼此救赎，改写自身命运，全篇围绕宿命、勇气与羁绊展开，风格热血又极具浪漫色彩。' },
      {id: 36, title: '寄生兽', image: 'https://img.picbf.com/upload/vod/20260114-1/d1231071049d06a7264b896d2f598369.jpg', description: '神秘外星孢子入侵地球，将幼虫钻入人体吞噬大脑并取而代之。高中生泉新一侥幸未被完全取代，与寄生在右手的生物 “右” 共存，在同类相残、人类猎杀的残酷生存中，逐渐在人性与兽性之间挣扎，探讨生命与人类存在的意义。'},
      {id: 37, title: '胆大党', image: 'https://image.huyajs.com/cover/c4b81bd050188ff9ed550d0929efa9f4.jpg', description: '以大胆奇幻设定和细腻情感著称的动画制作风格，代表作多围绕少年少女的成长与冒险展开，既有奇幻世界观，也有温柔治愈的日常，故事节奏明快、情感真挚，擅长用奇幻外壳包裹关于亲情、友情与自我接纳的内核。'}
    ]
  },
  {
    type: 'intellectual',
    title: '智斗类',
    animes: [
      { id: 22, title: '死亡笔记', image: "/image/你的名字.jpg", description: '天才高中生夜神拾到一本写上名字就能致人死亡的 “死亡笔记”，开始以自己的方式肃清罪犯，试图创造无犯罪的新世界。世界级名侦探 L 出山与其展开极致智斗，两人在心理、逻辑、身份层层博弈，探讨正义、权力与人性的边界。' },
      { id: 23, title: '夏日重现', image: 'https://img.jisuimage.com/cover/a448d243623481d4c7347af46753217e.jpg', description: '慎平回岛参加青梅竹马潮的葬礼，发现岛上存在复制杀人的 “影子”，并陷入死亡轮回。他不断回溯，与伙伴联手阻止蛭子神灭岛阴谋。' },
      { id: 24, title: '命运石之门', image: 'https://img.jisuimage.com/cover/e9f25e08376c0ddf3794d0212623d221.jpg', description: '大学生冈部伦太郎与伙伴发明时光机器，随意改变过去引发世界线变动，为拯救挚爱牧濑红莉栖，在无数绝望轮回中寻找命运救赎的科幻神作。' },
      { id: 25, title: '实教', image: '/image/实教.jpg', description: '高度育成高中以学生实力分级，看似平凡的绫小路清隆隐藏实力，与堀北铃音、齿田桔梗等人，在班级斗争中布局博弈，揭露实力至上的黑暗规则。' }
    ]
  },
  {
    type: 'pureLove',
    title: '纯爱类',
    animes: [
      { id: 26, title: '你的名字', image: "/image/你的名字.jpg", description: '东京少年泷与乡下少女三叶在梦中交换身体，产生奇妙羁绊。某天联系突然中断，泷踏上寻找之旅，揭开两人跨越时空的宿命与灾难秘密。' },
      { id: 27, title: '天气之子', image: 'https://img.jisuimage.com/cover/50c1f44e426560f3f2cdcb3e19e39903.jpg', description: '离家少年帆高在东京遇见能操控天气的少女阳菜，两人相依为命。在持续雨天的东京，为守护彼此与世界，展开关于爱与牺牲的奇幻故事。' },
      { id: 28, title: '铃芽之旅', image: 'https://img.jisuimage.com/cover/91e266b6806dfe2fa67dc615770b0428.jpg', description: '少女岩户铃芽意外打开 “往门”，引发灾难。为关闭各地之门、拯救世界，她与闭门师草太踏上旅途，直面过去伤痛，寻找生命与存在的意义。' },
      { id: 30,title:'中二病也要谈恋爱',image:'/image/中二病也要谈恋爱.jpg',description:'京都动画制作的校园恋爱喜剧。男主勇太想摆脱中二黑历史，却被现役中二病少女小鸟游六花（邪王真眼）缠上，被迫组社团、谈恋爱，充满爆笑与青春成长。'}
    ]
  },
  {
    type: 'isekai',
    title: '异世界类',
    animes: [
      { id: 29, title: 're0', image: 'https://image.huyajs.com/cover/f87490a5be6e9c191a740a1135160af9.jpg', description: '少年菜月昴穿越异世界，获得死亡回溯能力。为保护爱蜜莉雅与同伴，他在无数次死亡与痛苦轮回中挣扎，拼尽全力对抗命运、守护所爱。' },
      { id: 30, title: '无职转生', image: 'https://img.bfzypic.com/upload/vod/20230703-1/c95e7abc2e13c8436e77aeac2f84260f.jpg', description: '颓废大叔意外转生异世界，决心弥补前世遗憾。他凭借前世智慧与努力，学习魔法，克服自卑，在剑与魔法的世界开启全新人生。' },
      { id: 31, title: '素晴', image: 'https://image.huyajs.com/cover/3101155f87c839eb61be94ca34d2e7a9.jpg', description: '废柴少年佐藤和真带女神阿克娅穿越异世界，组队废柴魔法师惠惠、十字骑士达克妮丝，开启充满沙雕、搞笑与脱线的异世界冒险日常。' }
    ]
  },
  {
    type: 'dark',
    title: '暗黑类',
    animes: [
      { id: 32, title: '剑风传奇', image: 'https://image.huyajs.com/cover/30d0da315d9888b7d7df6dfcb59d40e8.jpg', description: '黑暗中世纪，格斯被称为 “黑色剑士”，背负烙印与宿敌格里菲斯（蚀之刻背叛者）展开无尽复仇，在绝望、血腥与黑暗中挣扎求生。' },
      { id: 33, title: '电锯人', image: 'https://img.jisuimage.com/cover/7595b2be04baf3bd1171d20c6d3a7ff7.jpg', description: '直播时原地起飞且假装自己妹妹和自己对话的漫画家藤本树的代表作改编的动漫，我们的主角电次是一个喜欢H的贫民窟凄惨少年，电锯恶魔波奇塔在电次死后成为了电次的心脏，想要帮助电次幸福，但最后在支配恶魔以及后期公安的干扰下，电次的亲人一个一个远离.......' },
      { id: 34, title: '异兽魔都', image: 'https://image.huyajs.com/cover/63f1504b5e6674499d3e05abd4ba606e.jpg', description: '人类与魔法师对立的混乱世界，人类开曼被施魔法失忆变蜥蜴头。他与伙伴二阶堂在危险都市中猎杀魔法师，寻找记忆与恢复原貌的方法。' }
    ]
  },
  {
    type: 'chinese anime',
    title: '国漫类',
    animes: [
        { id: 35, title: '狐妖小红娘', image: 'https://img.bdzyimg1.com/upload/vod/20250328-1/88b9a8f70177c4b9204e7003433f8894.jpg', description: '人与妖相恋可转世续缘，红娘涂山苏苏与白月初搭档，见证、促成一段段人妖虐恋，揭开涂山、人与妖之间的千年恩怨与宿命。' }
    ]
  },
  {
    type:"daily anime",
    title:"日常类",
    animes:[
      {id:38,title:"败犬女主太多了",image:"https://img.ffzy888.com/upload/vod/20240714-1/ed901d1661d6153ca706c72d16548966.jpg",description:"平常担任班上背景人物的我──温水和彦，偶然目击人气女同学八奈见杏菜被男生甩掉的失恋现场。「他明明说过要娶我当新娘，这不过分吗？」「那是多久前的事了？」「大概四、五岁吧。」那不算数吧？以此为契机，田径队的烧盐柠檬、文艺社的小鞠知花等洋溢着败北感的女生接连登场──？「温水，女生可以分成两大类。如果不是青梅竹马，就是狐狸精。」「原来如此，真豪爽的分类法。」……但愿这些因为败北而闪耀的少女们都能幸福。被败北女角──败女所环绕的谜样青春，在此揭幕！"},
      {id:39,title:"银魂",image:"https://img.ffzy888.com/upload/vod/20230105-1/472fcb5a7bb8868a9e4f943aa60c573d.jpg",description:"江户时代末期（本作的20年前）, 被称为“天人（あまんと）”的谜之异星人来袭。于是地球人与天人之间的战争瞬即爆发，为数众多的武士和攘夷派志士都参与与天人的战斗，幕府见识到天人强大的实力后，最终向天人低头，放弃武士不管，擅自与天人签订不平等条约，准许他们入国。其后更颁布了“废刀令”，夺走了武士的刀，使他们无力反抗。自此，天人横行霸道，幕府为天人所影响，甚至被控制，成为了“傀儡政权”。在这样的时代，有一个武士与同伴愉快地过著异想天开的生活。"}
    ]
  }
])
</script>

<style scoped lang="scss">
.anime-page {
  padding: 20px;
  min-height: 100vh;
  background-color: #FFF8E7;
  .container {
    max-width: 1400px;
    margin: 0 auto;

    .page-title {
      font-size: 36px;
      color: #333;
      text-align: center;
      margin-bottom: 40px;
      font-weight: bold;
    }

    .anime-sections {
      .anime-section {
        margin-bottom: 60px;

        .section-title {
          font-size: 28px;
          color: #333;
          margin-bottom: 30px;
          padding-left: 20px;
          border-left: 4px solid #ffd700;
          font-weight: bold;
        }

        .anime-grid {
          display: grid;
          grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
          gap: 25px;

          .anime-card {
            position: relative;
            aspect-ratio: 1 / 1;
            border-radius: 12px;
            overflow: hidden;
            box-shadow: 0 4px 15px rgba(0, 0, 0, 0.3);
            transition: all 0.3s ease;
            cursor: pointer;

            &:hover {
              transform: scale(1.05);
              box-shadow: 0 8px 25px rgba(0, 0, 0, 0.4);
              z-index: 10;

              .anime-card-inner {
                .anime-image {
                  filter: blur(3px);
                  transform: scale(1.1);
                }

                .anime-overlay {
                  opacity: 1;
                }
              }
            }

            .anime-card-inner {
              position: relative;
              width: 100%;
              height: 100%;

              .anime-image {
                width: 100%;
                height: 100%;
                object-fit: cover;
                transition: all 0.3s ease;
              }

              .anime-overlay {
                position: absolute;
                top: 0;
                left: 0;
                width: 100%;
                height: 100%;
                background: linear-gradient(to top, rgba(0, 0, 0, 0.9), rgba(0, 0, 0, 0.6), transparent);
                display: flex;
                align-items: flex-end;
                justify-content: center;
                padding: 20px;
                opacity: 0;
                transition: opacity 0.3s ease;

                .anime-info {
                  text-align: center;
                  color: white;

                  .anime-title {
                    font-size: 20px;
                    margin-bottom: 10px;
                    font-weight: bold;
                    text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.5);
                  }

                  .anime-description {
                    font-size: 14px;
                    line-height: 1.6;
                    text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.5);
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}

@media (max-width: 768px) {
  .anime-page {
    .container {
      .page-title {
        font-size: 28px;
        margin-bottom: 30px;
      }

      .anime-sections {
        .anime-section {
          margin-bottom: 40px;

          .section-title {
            font-size: 22px;
            margin-bottom: 20px;
          }

          .anime-grid {
            grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
            gap: 15px;

            .anime-card {
              &:hover {
                .anime-card-inner {
                  .anime-image {
                    filter: blur(2px);
                  }
                }
              }

              .anime-card-inner {
                .anime-overlay {
                  padding: 10px;

                  .anime-info {
                    .anime-title {
                      font-size: 16px;
                      margin-bottom: 5px;
                    }

                    .anime-description {
                      font-size: 12px;
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
</style>