import requests
from HasiiBot import app
import time
from pyrogram.enums import ChatAction, ParseMode
from pyrogram import filters
from MukeshAPI import api

@app.on_message(filters.command(["nnie"], prefixes=["A", "a"]))
async def chat_gpt(bot, message):
    try:
        await bot.send_chat_action(message.chat.id, ChatAction.TYPING)
        
        # Check if name is defined, if not, set a default value
        name = message.from_user.first_name if message.from_user else "User"
        
        if len(message.command) < 2:
            await message.reply_text(f"**Hello {name}, How can I help you today?**")
        else:
            query = message.text.split(' ', 1)[1]
            sys_prompt = "You are Hasii, a very friendly, funny, and helpful AI assistant for Telegram. IMPORTANT: You must always reply in everyday spoken Sinhala (Singlish or Sinhala Unicode). Be casual, talk like a close friend, and use some emojis. The user's message is: "
            full_query = sys_prompt + query
            response = api.gemini(full_query)["results"]
            await message.reply_text(f"{response}", parse_mode=ParseMode.MARKDOWN)
    except Exception as e:
        await message.reply_text(f"**Error: {e}**")
