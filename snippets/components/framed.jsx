export const Framed = ({
    image,
    background,
    link,
    alt = "Framed content",
    borderWidth = "24px",
    outerRadius = "12px",
    innerRadius = "10px",
    size = "100%",
    offsetX = "0px",
    offsetY = "0px",
}) => {
    const innerImage = (
        <img
            src={image}
            alt={alt}
            noZoom={link ? true : false}
            style={{
                display: "block",
                width: size,
                height: size,
                borderRadius: innerRadius,
                margin: `0 0`,
                transform: `translate(${offsetX}, ${offsetY})`,
            }}
        />
    );

    return (
        <div
            style={{
                backgroundImage: background ? `url(${background})` : undefined,
                backgroundSize: "cover",
                borderRadius: outerRadius,
                padding: `${borderWidth}`,
                overflow: "hidden",
            }}
        >
            <div>
                {link ? (<a href={link}> {innerImage} </a>) : (innerImage)}
            </div>
        </div>
    );
};
