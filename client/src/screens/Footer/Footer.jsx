import React from "react";
import "./footer.scss";
import logo from "../../assets/Footer/Logo.svg";
import { Link } from "react-router-dom";
const Footer = () => {
  return (
    <div className="footer">
      <div className="footer__container">
        <Link to="/main">
          <img src={logo} alt="" className="footer__logo" />
        </Link>
        <div className="footer__contacts">
          <p className="footer__contacts-title">Контакты:</p>
          <ul className="footer__contacts-list">
            <li>
              Telegram:{" "}
              <a
                href="https://t.me/Rade16"
                target="_blank"
                className="footer__contacts-list-links"
              >
                @Rade16
              </a>
            </li>
            <li>
              Vk:{" "}
              <a
                href="https://vk.com/clownyaara"
                target="_blank"
                className="footer__contacts-list-links"
              >
                @clownyaara
              </a>
            </li>
            <li>
              Почта:{" "}
              <a
                href="mailto:reportingjournal@gmail.com"
                target="_blank"
                className="footer__contacts-list-links"
              >
                reportingjournal@gmail.com
              </a>
            </li>
          </ul>
        </div>

        <div className="footer__other">
          <ul className="footer__other-list">
            <li>Пользовательское соглашение</li>
            <li>Справка</li>
          </ul>
        </div>
        <div className="footer__company">
          <p>© 2024 ООО «Reporting journal»</p>
          <p>Проект компании Мактрахер</p>
        </div>
      </div>
    </div>
  );
};

export default Footer;
