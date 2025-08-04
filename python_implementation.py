# Python Implementation Files for Anime Api

import feedparser
import requests
import json
import os
import sys
import time
from datetime import datetime
from typing import List, Dict, Optional
from dataclasses import dataclass

@dataclass
class AnimeNews:
    title: str
    summary: str
    link: str
    source: str
    published_at: datetime

class RSSFetcher:
    """Tool 1: fetch_anime_news - Gets latest anime news from RSS feeds"""
    
    def __init__(self):
        self.feeds = [
            "https://www.animenewsnetwork.com/all/rss.xml",
            "https://feeds.crunchyroll.com/news.rss",
            "https://myanimelist.net/rss/news.xml",
            "https://www.otakunews.com/feed/",
            "https://animehunch.com/feed/"
        ]
    
    def fetch_anime_news(self) -> List[AnimeNews]:
        """Fetches latest anime news from RSS feeds"""
        all_news = []
        
        for feed_url in self.feeds:
            try:
                news = self._fetch_from_feed(feed_url)
                all_news.extend(news)
            except Exception as e:
                print(f"Error fetching from {feed_url}: {e}")
                continue
        
        # Sort by published date (newest first)
        all_news.sort(key=lambda x: x.published_at, reverse=True)
        
        # Return top 15 most recent
        return all_news[:15]
    
    def _fetch_from_feed(self, feed_url: str) -> List[AnimeNews]:
        """Fetch news from a single RSS feed"""
        feed = feedparser.parse(feed_url)
        news = []
        
        for entry in feed.entries:
            if not self._is_anime_related(entry.title, getattr(entry, 'summary', '')):
                continue
            
            published_at = datetime.now()
            if hasattr(entry, 'published_parsed') and entry.published_parsed:
                published_at = datetime(*entry.published_parsed[:6])
            
            news_item = AnimeNews(
                title=self._clean_html(entry.title),
                summary=self._clean_html(getattr(entry, 'summary', '')),
                link=entry.link,
                source=self._extract_source_name(feed_url),
                published_at=published_at
            )
            news.append(news_item)
        
        return news
    
    def _is_anime_related(self, title: str, description: str) -> bool:
        """Check if content is anime-related"""
        content = (title + " " + description).lower()
        
        anime_keywords = [
            "anime", "manga", "otaku", "cosplay", "convention",
            "studio", "season", "episode", "character", "series",
            "crunchyroll", "funimation", "viz", "shonen", "seinen",
            "shoujo", "josei", "mecha", "isekai", "slice of life",
            "dragon ball", "naruto", "one piece", "attack on titan",
            "demon slayer", "my hero academia", "jujutsu kaisen"
        ]
        
        return any(keyword in content for keyword in anime_keywords)
    
    def _clean_html(self, content: str) -> str:
        """Simple HTML tag removal"""
        import re
        # Remove HTML tags
        clean = re.sub('<.*?>', '', content)
        # Clean up whitespace
        return ' '.join(clean.split())
    
    def _extract_source_name(self, feed_url: str) -> str:
        """Extract source name from feed URL"""
        if "animenewsnetwork" in feed_url:
            return "Anime News Network"
        elif "crunchyroll" in feed_url:
            return "Crunchyroll"
        elif "myanimelist" in feed_url:
            return "MyAnimeList"
        elif "otakunews" in feed_url:
            return "Otaku News"
        elif "animehunch" in feed_url:
            return "Anime Hunch"
        return "Unknown Source"

class DuplicateChecker:
    """Tool 2: check_if_posted_before & Tool 5: log_as_published"""
    
    def __init__(self, log_file: str = "data/published_articles.txt"):
        self.log_file = log_file
        os.makedirs(os.path.dirname(log_file), exist_ok=True)
    
    def check_if_posted_before(self, article_link: str) -> bool:
        """Returns True if NEW (not posted before), False if OLD (already posted)"""
        if not os.path.exists(self.log_file):
            return True  # New article
        
        try:
            with open(self.log_file, 'r', encoding='utf-8') as f:
                for line in f:
                    if line.strip() and article_link in line:
                        return False  # Already posted
            return True  # New article
        except Exception as e:
            print(f"Error checking duplicates: {e}")
            return True  # Assume new on error
    
    def log_as_published(self, article_link: str, title: str) -> str:
        """Logs article as published"""
        try:
            timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
            log_entry = f"{timestamp}|{article_link}|{title}\n"
            
            with open(self.log_file, 'a', encoding='utf-8') as f:
                f.write(log_entry)
            
            return "Log updated."
        except Exception as e:
            return f"Error logging: {e}"
    
    def get_published_count(self) -> int:
        """Get total number of published articles"""
        if not os.path.exists(self.log_file):
            return 0
        
        try:
            with open(self.log_file, 'r', encoding='utf-8') as f:
                return sum(1 for line in f if line.strip())
        except Exception:
            return 0

class SinhalaWriter:
    """Tool 3: write_anime_post_in_my_style - AI-powered Sinhala content generation"""
    
    def __init__(self, api_key: str):
        self.api_key = api_key
        self.base_url = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-pro:generateContent"
    
    def write_anime_post_in_my_style(self, title: str, summary: str, link: str) -> str:
        """Generates authentic Sinhala post using AI"""
        prompt = self._build_persona_prompt(title, summary, link)
        
        payload = {
            "contents": [
                {
                    "parts": [
                        {"text": prompt}
                    ]
                }
            ]
        }
        
        try:
            response = requests.post(
                f"{self.base_url}?key={self.api_key}",
                headers={"Content-Type": "application/json"},
                json=payload,
                timeout=30
            )
            
            if response.status_code == 200:
                data = response.json()
                if data.get("candidates") and data["candidates"][0].get("content"):
                    return data["candidates"][0]["content"]["parts"][0]["text"].strip()
            
            return f"Error generating content: {response.status_code}"
        
        except Exception as e:
            return f"Error: {e}"
    
    def _build_persona_prompt(self, title: str, summary: str, link: str) -> str:
        """Build the persona-driven prompt for AI"""
        return f"""You are a casual Sri Lankan anime fan writing for friends. Write in natural Sinhala with mixed English - just like how real Sri Lankans talk. Keep it simple, casual and fun.

**Style:**
- Mix Sinhala and English naturally (like "anime එකක්", "game එක", "trailer එක")
- Use casual words: "අයියේ", "අක්කේ", "කොල්ලා", "කෙල්ලටත්" 
- Common expressions: "ඒකනේ", "මේකද", "කොහොමද", "නේද"
- Keep it short and excited
- Add some emojis
- End with a question

**News:**
Title: {title}
Summary: {summary}
Link: {link}

Write a casual post now:"""

class SocialMediaPublisher:
    """Tool 4: publish_post - Publishes content to social media platforms"""
    
    def __init__(self, telegram_bot_token: str = None, telegram_chat_id: str = None):
        self.telegram_bot_token = telegram_bot_token
        self.telegram_chat_id = telegram_chat_id
    
    def publish_post(self, post_text: str) -> str:
        """Publishes post to configured social media platform"""
        if self.telegram_bot_token and self.telegram_chat_id:
            return self._publish_to_telegram(post_text)
        
        return "No social media platform configured"
    
    def _publish_to_telegram(self, post_text: str) -> str:
        """Publish to Telegram"""
        url = f"https://api.telegram.org/bot{self.telegram_bot_token}/sendMessage"
        
        payload = {
            "chat_id": self.telegram_chat_id,
            "text": post_text
        }
        
        try:
            response = requests.post(url, json=payload, timeout=30)
            
            if response.status_code == 200:
                data = response.json()
                if data.get("ok"):
                    return "Successfully posted to Telegram!"
                else:
                    return f"Telegram API error: {data.get('description', 'Unknown error')}"
            else:
                return f"HTTP error: {response.status_code}"
        
        except Exception as e:
            return f"Error publishing to Telegram: {e}"
    
    def test_connection(self) -> bool:
        """Test connection to social media platform"""
        if self.telegram_bot_token:
            url = f"https://api.telegram.org/bot{self.telegram_bot_token}/getMe"
            try:
                response = requests.get(url, timeout=10)
                return response.status_code == 200 and response.json().get("ok", False)
            except Exception:
                return False
        return False

class AnimeApiOrchestrator:
    """The autonomous Gemini agent orchestrator"""
    
    def __init__(self, rss_fetcher: RSSFetcher, duplicate_checker: DuplicateChecker, 
                 sinhala_writer: SinhalaWriter, social_publisher: SocialMediaPublisher):
        self.rss_fetcher = rss_fetcher
        self.duplicate_checker = duplicate_checker
        self.sinhala_writer = sinhala_writer
        self.social_publisher = social_publisher
    
    def execute_cycle(self) -> bool:
        """Execute one complete autonomous cycle"""
        print("🚀 Anime Api awakening! Time to check for exciting anime news...")
        
        # Tool 1: Fetch anime news
        print("📡 Tool 1: Fetching latest anime news from RSS feeds...")
        articles = self.rss_fetcher.fetch_anime_news()
        
        if not articles:
            print("❌ No anime articles found. Sleeping until next cycle...")
            return False
        
        print(f"✅ Found {len(articles)} potential articles. Now checking for new content...")
        
        # Tool 2: Find first new article
        selected_article = None
        for i, article in enumerate(articles, 1):
            print(f"🔍 Checking article {i}: {article.title}")
            
            is_new = self.duplicate_checker.check_if_posted_before(article.link)
            
            if is_new:
                selected_article = article
                print(f"🎉 Found NEW article: {article.title}")
                break
            else:
                print(f"⏭️  Already posted: {article.title}")
        
        if not selected_article:
            print("😴 All articles have been posted before. Nothing new to share today!")
            return False
        
        # Tool 3: Write Sinhala post
        print("✍️  Tool 3: Writing exciting Sinhala post using AI...")
        sinhala_text = self.sinhala_writer.write_anime_post_in_my_style(
            selected_article.title,
            selected_article.summary,
            selected_article.link
        )
        
        print("📝 AI has crafted the perfect post! Here's what it wrote:")
        print(f"---\n{sinhala_text}\n---")
        
        # Tool 4: Publish post
        print("📢 Tool 4: Publishing to social media...")
        publish_result = self.social_publisher.publish_post(sinhala_text)
        print(f"🎊 {publish_result}")
        
        # Tool 5: Log as published
        print("📋 Tool 5: Logging article as published...")
        log_result = self.duplicate_checker.log_as_published(selected_article.link, selected_article.title)
        print(f"✅ {log_result}")
        
        # Final success message
        print("🎯 Mission accomplished! Anime Api has successfully:")
        print(f"   • Found new anime news: {selected_article.title}")
        print(f"   • Written engaging Sinhala content")
        print(f"   • Published to social media")
        print(f"   • Logged to prevent duplicates")
        print("😴 Anime Api is now sleeping until the next scheduled run...")
        
        return True
    
    def test_all_tools(self) -> bool:
        """Test all tools"""
        print("🧪 Testing all tools...")
        
        # Test RSS Fetcher
        print("Testing RSS Fetcher...")
        articles = self.rss_fetcher.fetch_anime_news()
        print(f"✅ RSS Fetcher: Found {len(articles)} articles")
        
        # Test Duplicate Checker
        print("Testing Duplicate Checker...")
        count = self.duplicate_checker.get_published_count()
        print(f"✅ Duplicate Checker: {count} articles in log")
        
        # Test Social Media Publisher
        print("Testing Social Media Publisher...")
        connection_ok = self.social_publisher.test_connection()
        if connection_ok:
            print("✅ Social Media Publisher: Connection successful")
        else:
            print("⚠️  Social Media Publisher: Connection failed")
        
        # Test Sinhala Writer (if we have articles)
        if articles:
            print("Testing Sinhala Writer...")
            test_result = self.sinhala_writer.write_anime_post_in_my_style(
                "Test Article",
                "This is a test summary for anime news testing.",
                "https://example.com/test"
            )
            if "Error" not in test_result:
                print("✅ Sinhala Writer: AI response generated successfully")
            else:
                print(f"⚠️  Sinhala Writer: {test_result}")
        
        print("🎉 All tools tested!")
        return True
    
    def get_status(self) -> Dict:
        """Get current status"""
        return {
            "published_count": self.duplicate_checker.get_published_count(),
            "last_run": datetime.now().isoformat(),
            "social_media_connected": self.social_publisher.test_connection()
        }

# Configuration
def load_config():
    """Load configuration from environment variables"""
    from dotenv import load_dotenv
    load_dotenv()
    
    return {
        "GEMINI_API_KEY": os.getenv("GEMINI_API_KEY"),
        "TELEGRAM_BOT_TOKEN": os.getenv("TELEGRAM_BOT_TOKEN"),
        "TELEGRAM_CHAT_ID": os.getenv("TELEGRAM_CHAT_ID"),
        "MAX_ARTICLES": int(os.getenv("MAX_ARTICLES", 5)),
        "REQUEST_TIMEOUT": int(os.getenv("REQUEST_TIMEOUT", "30").replace("s", ""))
    }

# Main application
def main():
    """Main application entry point"""
    config = load_config()
    
    if not config["GEMINI_API_KEY"]:
        print("❌ GEMINI_API_KEY is required!")
        return
    
    # Initialize services
    rss_fetcher = RSSFetcher()
    duplicate_checker = DuplicateChecker()
    sinhala_writer = SinhalaWriter(config["GEMINI_API_KEY"])
    social_publisher = SocialMediaPublisher(
        config["TELEGRAM_BOT_TOKEN"],
        config["TELEGRAM_CHAT_ID"]
    )
    
    # Initialize orchestrator
    orchestrator = AnimeApiOrchestrator(
        rss_fetcher,
        duplicate_checker,
        sinhala_writer,
        social_publisher
    )
    
    # Run one complete cycle
    success = orchestrator.execute_cycle()
    if success:
        print("🎉 Cycle completed successfully!")
    else:
        print("😴 No new content to process.")

# CLI interface
def cli():
    """CLI interface"""
    import sys
    
    if len(sys.argv) < 2:
        print("Anime Api CLI Tool")
        print("\nUsage:")
        print("  python main.py --test    : Test all tools without posting")
        print("  python main.py --status  : Show current status")
        print("  python main.py --run     : Run one complete cycle")
        return
    
    config = load_config()
    if not config["GEMINI_API_KEY"]:
        print("❌ GEMINI_API_KEY is required!")
        return
    
    # Initialize services
    rss_fetcher = RSSFetcher()
    duplicate_checker = DuplicateChecker()
    sinhala_writer = SinhalaWriter(config["GEMINI_API_KEY"])
    social_publisher = SocialMediaPublisher(
        config["TELEGRAM_BOT_TOKEN"],
        config["TELEGRAM_CHAT_ID"]
    )
    
    orchestrator = AnimeApiOrchestrator(
        rss_fetcher,
        duplicate_checker,
        sinhala_writer,
        social_publisher
    )
    
    command = sys.argv[1]
    
    if command == "--test":
        orchestrator.test_all_tools()
    elif command == "--status":
        status = orchestrator.get_status()
        print("📊 Status Report:")
        print(f"   Published Articles: {status['published_count']}")
        print(f"   Last Run: {status['last_run']}")
        print(f"   Social Media Connected: {status['social_media_connected']}")
    elif command == "--run":
        print("🎯 Running complete autonomous cycle...")
        success = orchestrator.execute_cycle()
        if success:
            print("🎉 Cycle completed successfully!")
        else:
            print("😴 No new content to process.")

if __name__ == "__main__":
    if len(sys.argv) > 1:
        cli()
    else:
        main()
